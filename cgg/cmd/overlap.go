package cmd

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/j6k4m8/cg/cgg/runner"
)

type read struct {
	Label string `json:"label"`
	Seq   string `json:"-"`
}

type readPair struct {
	Left    *read `json:"left"`
	Right   *read `json:"right"`
	Overlap int   `json:"overlap"`
}

func Overlap(context *cli.Context) *Response {
	path := context.Args().First()
	if path == "" {
		return ErrorMissingArgument()
	}

	pairs, err := computeOverlap(path)

	if err != nil {
		return ErrorOccured(err)
	}

	// pairLabels := make(map[string]string)
	// for k, v := range pairs {
	// 	pairLabels[k] = v.Right.Label
	// }

	return &Response{
		Ok:      true,
		Content: pairs,
	}

}

func computeOverlap(path string) (map[string]readPair, error) {

	reads, err := parseFasta(path)
	if err != nil {
		return nil, err
	}

	r := runner.NewMax(
		func(i int, r runner.Runner) (interface{}, error) {
			step := int(len(reads) / r.NumRoutines())
			start := step * i
			end := start + step
			if end > len(reads) {
				end = len(reads)
			}
			return findOverlaps(reads, start, end)
		},
	)

	pairs := make(map[string]readPair)
	r.Run()
	results, _ := r.Wait()

	for _, r := range results {
		rS := r.([]readPair)
		for _, rP := range rS {
			pairs[rP.Left.Label] = rP
		}
	}

	return pairs, nil
}

func parseFasta(path string) ([]read, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	var reads []read

	var seqBuf bytes.Buffer
	var name string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if seqBuf.Len() > 0 {
				reads = append(reads, read{Label: name, Seq: seqBuf.String()})
				seqBuf.Reset()
			}
			name = strings.TrimSpace(line[1:])
		} else {
			seqBuf.WriteString(strings.TrimSpace(line))
		}
	}
	if seqBuf.Len() > 0 {
		reads = append(reads, read{Label: name, Seq: seqBuf.String()})
	}

	return reads, nil
}

func findOverlaps(reads []read, start, end int) ([]readPair, error) {
	ret := make([]readPair, 0, end-start)
	for i := start; i < end; i++ {
		tR := reads[i]
		minLen := 40
		minLenRead := -1
		for j, pR := range reads {
			if i == j {
				continue
			}
			lenOverlap := suffixPrefixMatch(tR.Seq, pR.Seq, minLen)
			if lenOverlap > minLen {
				minLen = lenOverlap
				minLenRead = j
			} else if lenOverlap == minLen {
				minLenRead = -1
			}
		}
		if minLenRead != -1 {
			ret = append(ret, readPair{
				Left: &tR, Right: &reads[minLenRead], Overlap: minLen,
			})
		}
	}
	return ret, nil
}

func suffixPrefixMatch(str1, str2 string, minOverlap int) int {
	if len(str1) < minOverlap {
		return 0
	}
	str2Prefix := str2[:minOverlap]
	str1Pos := -1
	for {
		str1Pos = index(str1, str2Prefix, str1Pos+1)
		if str1Pos == -1 {
			return 0
		}
		str1Suffix := str1[str1Pos:]
		if strings.HasPrefix(str2, str1Suffix) {
			return len(str1Suffix)
		}
	}
	return -1
}

func index(str1, str2 string, start int) int {
	pos := strings.Index(str1[start:], str2)
	if pos == -1 {
		return -1
	}
	return pos + start
}
