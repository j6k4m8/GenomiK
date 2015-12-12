package cmd

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/j6k4m8/cg/cgg/runner"
)

const SEP = ","

type unitigSorter []*fullUnitig

func (u unitigSorter) Len() int {
	return len(u)
}

func (u unitigSorter) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u unitigSorter) Less(i, j int) bool {
	return u[i].Reads[0].Left.Label < u[j].Reads[0].Left.Label
}

type fullUnitig struct {
	*unitig
	Seq string `json:"seq"`
}

func minTwo(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Assemble(context *cli.Context) *Response {
	path := context.Args().First()
	if path == "" {
		return ErrorMissingArgument()
	}

	partUnitigs, err := computeUnitigs(path, context.Bool(GZipFlag))
	if err != nil {
		return ErrorOccured(err)
	}

	unitigs := completeUnitigs(partUnitigs)

	// for testing purposes we can sort them to guarentee order
	// sort.Sort(unitigSorter(unitigs))

	// compute the matrix concurrently - very similar to how we compute
	// overlaps
	r := runner.New(
		// given an anonymous function so we can pass the local unitigs slice
		func(i int, r runner.Runner) (interface{}, error) {
			return computeMatrix(i, r, unitigs)
		},
		// run on either max procs OR the number of unitigs in the event that
		// is less than max procs
		minTwo(len(unitigs), runtime.GOMAXPROCS(-1)),
	)

	// run the computation
	r.Run()
	// wait for the results
	results, _ := r.Wait()

	// collect all results into one matrix
	// note that this depends on the runner keeping the correct order
	matrix := make([][]int, len(unitigs))
	pos := 0
	for _, outterA := range results {
		for _, innerA := range outterA.([][]int) {
			matrix[pos] = innerA
			pos++
		}
	}

	return &Response{
		Ok:      true,
		Content: matrix,
	}

}

func completeUnitigs(unitigs []*unitig) []*fullUnitig {
	fullU := make([]*fullUnitig, len(unitigs))
	for i := range unitigs {
		u := unitigs[i]
		var buf bytes.Buffer
		buf.WriteString(u.Reads[0].Left.Seq)
		for _, r := range u.Reads {
			buf.WriteString(r.Right.Seq[r.Overlap:])
		}
		fullU[i] = &fullUnitig{u, buf.String()}
	}
	return fullU
}

func computeMatrix(i int, r runner.Runner, unitigs []*fullUnitig) ([][]int, error) {
	step := int(len(unitigs) / r.NumRoutines())
	start := step * i
	end := start + step
	if end > len(unitigs) {
		end = len(unitigs)
	}
	results := make([][]int, end-start)

	for i := range results {
		results[i] = make([]int, len(unitigs))
	}
	for i := start; i < end; i++ {
		for j := range unitigs {
			if i == j {
				results[i-start][j] = -1
				continue
			}
			results[i-start][j] = suffixPrefixMatch(unitigs[i].Seq, unitigs[j].Seq, 0)
		}
	}

	return results, nil
}

func outputUnitigs(partUnitigs []*unitig, out io.Writer) error {
	for i := range partUnitigs {
		u := partUnitigs[i]
		err := outStr(u.Reads[0].Left.Seq, out)
		if err != nil {
			return err
		}
		for _, r := range u.Reads {
			if err := outStr(r.Right.Seq[r.Overlap:], out); err != nil {
				return err
			}
		}
		if i != len(partUnitigs)-1 {
			err = outStr(SEP, out)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func outStr(str string, out io.Writer) error {
	n, err := out.Write([]byte(str))
	if err != nil {
		return err
	} else if n < len(str) {
		return errors.New("Short write :-(")
	}
	return nil
}

type Flusher interface {
	Flush() error
}

type BufWriteCloser struct {
	flusher    Flusher
	fileCloser io.Closer
	stream     io.WriteCloser
}

func (b *BufWriteCloser) Close() error {
	// first close stream
	err := b.stream.Close()
	if err != nil {
		return err
	}
	// now flush buffer
	err = b.flusher.Flush()
	if err != nil {
		return err
	}
	// now close underlying file
	err = b.fileCloser.Close()
	if err != nil {
		return err
	}
	return nil
}

func (b *BufWriteCloser) Write(s []byte) (int, error) {
	return b.stream.Write(s)
}

type NopWriteCloser struct {
	io.Writer
}

func (n *NopWriteCloser) Close() error {
	return nil
}

func openWriter(path string, plaintext bool) (io.WriteCloser, error) {
	if !strings.HasSuffix(path, ".gz") {
		path += ".gz"
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	flusher := bufio.NewWriter(f)
	ret := &BufWriteCloser{flusher: flusher, fileCloser: f}
	if plaintext {
		ret.stream = &NopWriteCloser{f}
	} else {
		ret.stream = gzip.NewWriter(f)
	}
	return ret, nil
}
