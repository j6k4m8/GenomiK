package cmd

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/j6k4m8/cg/cgg/runner"
)

const GZipFlag = "gz"

// read holds the information for a single read from a FASTA file.
type read struct {
	Label string `json:"label"`
	Seq   string `json:"-"`
}

// readPair holds a pair of reads along with their computed overlap.
// This is ready to be marshalled by encoding/json.
type readPair struct {
	Left    *read `json:"left"`
	Right   *read `json:"right"`
	Overlap int   `json:"overlap"`
}

// Overlap provides a cmd handler for computing the overlaps (bbr) for a file
// containing FASTA reads.
//
// On success, this will return a *Response object containing all overlap pairs
// in the form of map of LeftLabel --> *readPair.
func Overlap(context *cli.Context) *Response {
	path := context.Args().First()
	if path == "" {
		return ErrorMissingArgument()
	}

	pairs, err := computeOverlap(path, context.Bool(GZipFlag))

	if err != nil {
		return ErrorOccured(err)
	}

	return &Response{
		Ok:      true,
		Content: pairs,
	}

}

// computeOverlap is the internal function that does all the heavy lifting of
// the Overlap command handler. This is intended to be used when overlaps must
// be computed as a part of a larger algorithm.
func computeOverlap(path string, isGzipped bool) (map[string]readPair, error) {
	// read in fasta file
	reads, err := parseFasta(path, isGzipped)
	if err != nil {
		return nil, err
	}

	// run concurrently!
	r := runner.NewMax(
		// pass an anonymous function so we can use the local reads variable
		func(i int, r runner.Runner) (interface{}, error) {
			return findOverlaps(i, r, reads)
		},
	)

	// create a map to "sum" the result of the computations
	pairs := make(map[string]readPair)
	// run the overlap functions!
	r.Run()
	// wait for and collect the results
	results, _ := r.Wait()

	for _, r := range results {
		// cast generic slice back to original type
		rS := r.([]readPair)
		for _, rP := range rS {
			pairs[rP.Left.Label] = rP
		}
	}

	return pairs, nil
}

// parseFasta takes a path string, opens the file, and parses it into the FASTA
// reads that it contains.
func parseFasta(path string, isGzipped bool) ([]read, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var unbufReader io.ReadCloser = file
	if isGzipped {
		unbufReader, err = gzip.NewReader(unbufReader)
		if err != nil {
			return nil, err
		}
		defer unbufReader.Close()
	}

	reader := bufio.NewReader(unbufReader)
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

func findOverlaps(i int, r runner.Runner, reads []read) ([]readPair, error) {

	step := int(len(reads) / r.NumRoutines())
	start := step * i
	end := start + step
	if end > len(reads) {
		end = len(reads)
	}
	var ret []readPair
	for i := start; i < end; i++ {
		tR := &reads[i]
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
				Left: tR, Right: &reads[minLenRead], Overlap: minLen,
			})
		}
		// maxLen := 40
		// var maxRead *read
		// matches, exists := prefixMap[tR.Seq[len(tR.Seq)-40:]]
		// if !exists {
		// 	continue
		// }
		// for j := range matches {
		// 	pR := &reads[j]
		// 	if tR.Label == pR.Label {
		// 		continue
		// 	}
		// 	lenO := suffixPrefixMatch(tR.Seq, pR.Seq, maxLen)
		// 	if lenO > maxLen {
		// 		maxLen = lenO
		// 		maxRead = pR
		// 	} else if lenO == maxLen {
		// 		maxRead = nil
		// 	}
		// }
		// if maxRead != nil {
		// 	ret = append(ret, readPair{
		// 		Left: tR, Right: maxRead, Overlap: maxLen,
		// 	})
		// }
	}

	return ret, nil
}

func suffixPrefixMatch(str1, str2 string, minOverlap int) int {
	if len(str1) < minOverlap || len(str2) < minOverlap {
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
