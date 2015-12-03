package cmd

import "github.com/codegangsta/cli"

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

	unitigs, err := computeUnitigs(path)
	if err != nil {
		return ErrorOccured(err)
	}

	uStart := make(map[string]int, 0)
	for k, v := range unitigs {
		uStart[k] = len(v.Reads)
	}

	return &Response{
		Ok:      true,
		Content: uStart,
	}
}

func computeUnitigs(path string) (map[string]*unitig, error) {
	tMap, err := computeOverlap(path)
	if err != nil {
		return nil, err
	}

	tLabels := make([]string, len(tMap))
	for k := range tMap {
		tLabels = append(tLabels, k)
	}

	pMap := parseOverlaps(tMap)

	unitigs := findUnitigs(tMap, pMap)
	return unitigs, nil
}

func findUnitigs(tMap map[string]readPair, pMap map[string]int) map[string]*unitig {
	uMap := make(map[string]*unitig)
	used := NewStringSet()
	for tLabel, pair := range tMap {
		if !used.AddContains(tLabel) {
			continue
		}
		pLabel := pair.Right.Label
		if pMap[pLabel] != pair.Overlap {
			continue
		}
		currU := newUnitig(pair)
		uMap[tLabel] = currU
		if uP, exists := uMap[pLabel]; exists {
			currU.addUnitig(uP)
			delete(uMap, pLabel)
			continue
		}
		for {
			tLabelNew := pLabel
			used.Add(tLabelNew)
			pairNew, exists := tMap[tLabelNew]
			if !exists {
				break
			}
			pLabel = pairNew.Right.Label
			if pMap[pLabel] != pairNew.Overlap {
				break
			}
			currU.add(pairNew)
			if u, exists := uMap[pLabel]; exists {
				currU.addUnitig(u)
				delete(uMap, pLabel)
				break
			}
		}
	}

	return uMap
}

// def find_unitigs(tMap, pMap):
//     unitigs = dict()
//     used = set()
//     for t in tMap.keys():
//         if t in used:
//             continue
//         used.add(t)
//         p_pair = tMap[t]
//         if pMap[p_pair.label] != p_pair.l:
//             continue
//         unitigs[t] = [p_pair]
//         if p_pair.label in unitigs:
//             unitigs[t].extend(unitigs[p_pair.label])
//             del unitigs[p_pair.label]
//             continue

//         while True:
//             t_new = p_pair.label
//             used.add(t_new)
//             if t_new not in tMap:
//                 break
//             p_pair = tMap[t_new]
//             if pMap[p_pair.label] != p_pair.l:
//                 break
//             unitigs[t].append(p_pair)
//             if p_pair.label in unitigs:
//                 unitigs[t].extend(unitigs[p_pair.label])
//                 del unitigs[p_pair.label]
//                 break
//     return unitigs

func parseOverlaps(tMap map[string]readPair) map[string]int {
	pMap := make(map[string]int)

	for _, pair := range tMap {
		pLabel := pair.Right.Label
		if _, exists := pMap[pLabel]; !exists {
			pMap[pLabel] = pair.Overlap
		} else if abs(pMap[pLabel]) == pair.Overlap {
			if pMap[pLabel] > 0 {
				pMap[pLabel] *= -1
			}
		} else if abs(pMap[pLabel]) < pair.Overlap {
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
