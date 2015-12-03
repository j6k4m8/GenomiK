package cmd

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
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

	// get number of procs and spawn that many goroutines
	num := runtime.GOMAXPROCS(-1)

	// channels on channels on channels
	done := make(chan bool, num)
	out := make(chan readPair, num*2)

	step := int(len(reads) / num)
	for start := 0; start < len(reads); start += step {
		end := start + step
		if end > len(reads) {
			end = len(reads)
		}
		go findOverlaps(out, done, reads, start, end)
	}
	pairs := make(map[string]readPair)
	numDone := 0
	// for == while because go is go
	for numDone < num {
		// pick between channel output depending on availability
		// otherwise would be blocking and that's icky
		select {
		case <-done:
			numDone++
		case i := <-out:
			pairs[i.Left.Label] = i
		}
	}
	close(done)
	close(out)

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

func findOverlaps(out chan<- readPair, done chan<- bool, reads []read, start, end int) {
	defer func() { done <- true }()
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
			out <- readPair{Left: &tR, Right: &reads[minLenRead], Overlap: minLen}
		}
	}
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
