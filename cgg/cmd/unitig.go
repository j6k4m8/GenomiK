package cmd

import "github.com/codegangsta/cli"

const (
	// OutputFlag is the flag for requesting full unitig output.
	OutputFlag = "o"
	// PlainTextFlag is the boolean flag for plaintext output.
	PlainTextFlag = "t"
)

type unitig struct {
	Reads []readPair `json:"reads"`
}

func (u *unitig) add(p readPair) {
	u.Reads = append(u.Reads, p)
}

func (u *unitig) addUnitig(uO *unitig) {
	for _, r := range uO.Reads {
		u.Reads = append(u.Reads, r)
	}
}

func newUnitig(p readPair) *unitig {
	uA := make([]readPair, 1)
	uA[0] = p
	return &unitig{Reads: uA}
}

func Unitig(context *cli.Context) *Response {
	path := context.Args().First()
	if path == "" {
		return ErrorMissingArgument()
	}
	outputPath := context.String(OutputFlag)

	unitigs, err := computeUnitigs(path, context.Bool(GZipFlag), context)
	if err != nil {
		return ErrorOccured(err)
	}

	uStart := make(map[string]int)
	for _, v := range unitigs {
		uStart[v.Reads[0].Left.Label] = len(v.Reads)
	}

	if len(outputPath) > 0 {
		out, err := openWriter(outputPath, context.Bool(PlainTextFlag))
		if err != nil {
			return ErrorOccured(err)
		}
		defer out.Close()
		err = outputUnitigs(unitigs, out)
		if err != nil {
			return ErrorOccured(err)
		}
	}

	return &Response{
		Ok:      true,
		Content: uStart,
	}
}

func computeUnitigs(path string, isGzipped bool, context *cli.Context) ([]*unitig, error) {
	tMap, err := computeOverlap(path, isGzipped, context)
	if err != nil {
		return nil, err
	}

	pMap := parseOverlaps(tMap)

	unitigs := traverseUnitigs(findMinUnitigs(tMap, pMap))
	return unitigs, nil
}

func findMinUnitigs(tMap map[string]readPair, pMap map[string]int) map[string]*unitig {
	unitigs := make(map[string]*unitig)
	for tLabel, pair := range tMap {
		if pair.Overlap == pMap[pair.Right.Label] {
			unitigs[tLabel] = newUnitig(pair)
		}
	}
	return unitigs
}

func traverseUnitigs(unitigs map[string]*unitig) []*unitig {
	var realUnitigs []*unitig
	revUni := make(map[string]string)
	uKeySlice := make([]string, 0, len(unitigs))
	for k, v := range unitigs {
		revUni[v.Reads[0].Right.Label] = k
		uKeySlice = append(uKeySlice, v.Reads[0].Left.Label)
	}
	uKeys := NewStringSetFromSlice(uKeySlice)
	uKeySlice = nil
	for !uKeys.IsEmpty() {
		start := uKeys.Pop()
		// traverse until we find the true start of the full unitig
		for {
			newStart, exists := revUni[start]
			if !exists {
				break
			}
			start = newStart
			uKeys.Remove(start)
		}

		u := newUnitig(unitigs[start].Reads[0])
		realUnitigs = append(realUnitigs, u)

		curr := start
		for {
			uKeys.Remove(curr)
			oldUni, exists := unitigs[curr]
			if !exists {
				break
			}
			u.addUnitig(oldUni)
			curr = oldUni.Reads[0].Right.Label
		}
	}
	return realUnitigs
}

func parseOverlaps(tMap map[string]readPair) map[string]int {
	pMap := make(map[string]int)

	for _, pair := range tMap {
		pLabel := pair.Right.Label
		if oldVal, exists := pMap[pLabel]; !exists {
			pMap[pLabel] = pair.Overlap
		} else if abs(oldVal) == pair.Overlap {
			if oldVal > 0 {
				pMap[pLabel] *= -1
			}
		} else if abs(oldVal) < pair.Overlap {
			pMap[pLabel] = pair.Overlap
		}
	}

	return pMap
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
