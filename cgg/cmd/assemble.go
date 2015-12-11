package cmd

import (
	"bytes"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/j6k4m8/cg/cgg/runner"
)

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

	partUnitigs, err := computeUnitigs(path)
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
