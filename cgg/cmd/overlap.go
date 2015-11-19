package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
)

type read struct {
	label,
	seq string
}

func Overlap(context *cli.Context) *Response {
	path := context.Args().First()
	if path == "" {
		return ErrorMissingArgument()
	}

	reads, err := parse_fasta(path)
	if err != nil {
		return ErrorOccured(err)
	}

	// get number of procs and spawn that many goroutines
	num := runtime.GOMAXPROCS(-1)

	// channels on channels on channels
	done := make(chan bool, num)
	out := make(chan string, num*2)

	step := int(len(reads) / num)
	for start := 0; start < len(reads); start += step {
		end := start + step
		if end > len(reads) {
			end = len(reads)
		}
		go find_overlaps(out, done, reads, start, end)
	}
	var buf bytes.Buffer
	num_done := 0
	// for == while because go is go
	for num_done < num {
		// pick between channel output depending on availability
		// otherwise would be blocking and that's icky
		select {
		case <-done:
			num_done++
		case i := <-out:
			buf.WriteString(i)
			buf.WriteByte('\n')
		}
	}
	close(done)
	close(out)
	return &Response{
		Ok:      true,
		Content: buf.String(),
	}

}

func parse_fasta(path string) ([]read, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	reads := make([]read, 0)

	var seqBuf bytes.Buffer
	var name string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if seqBuf.Len() > 0 {
				reads = append(reads, read{label: name, seq: seqBuf.String()})
				seqBuf.Reset()
			}
			name = strings.TrimSpace(line[1:])
		} else {
			seqBuf.WriteString(strings.TrimSpace(line))
		}
	}
	if seqBuf.Len() > 0 {
		reads = append(reads, read{label: name, seq: seqBuf.String()})
	}

	return reads, nil
}

func find_overlaps(out chan<- string, done chan<- bool, reads []read, start, end int) {
	defer func() { done <- true }()
	for i := start; i < end; i++ {
		t_r := reads[i]
		min_len := 40
		min_len_label := ""
		for j, p_r := range reads {
			if i == j {
				continue
			}
			len_overlap := suffix_prefix_match(t_r.seq, p_r.seq, min_len)
			if len_overlap > min_len {
				min_len = len_overlap
				min_len_label = p_r.label
			} else if len_overlap == min_len {
				min_len_label = ""
			}
		}
		if min_len_label != "" {
			out <- fmt.Sprintf("%s %s %d", t_r.label, min_len_label, min_len)
		}
	}
}

func suffix_prefix_match(str1, str2 string, min_overlap int) int {
	if len(str1) < min_overlap {
		return 0
	}
	str2_prefix := str2[:min_overlap]
	str1_pos := -1
	for {
		str1_pos = index(str1, str2_prefix, str1_pos+1)
		if str1_pos == -1 {
			return 0
		}
		str1_suffix := str1[str1_pos:]
		if strings.HasPrefix(str2, str1_suffix) {
			return len(str1_suffix)
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
