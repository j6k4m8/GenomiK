package cmd

// An alignment stores information about the actual string as well
// as its position.
type alignment struct {
	offset int
	seq    string
}

// Calculate the "cost" (for insertion into alignment matrices) of two strings
// @j6k4m8
func cost(c0, c1 byte) int {
	// Using the scores described by Smith et al (1981). This is effectively
	// changed to { 0, 2 } in the SW implementation because we max against 0.
	if c0 == c1 {
		return 2
	} else {
		return -1
	}
}

func ConcurrentSW(p string, t string) {

	h := make([][]int, len(p))
	for i := range h {
		h[i] = make([]int, len(t))
	}

	populate := func(i int, j int) {
		if i == 0 || j == 0 {
			h[i][j] = 0
		} else {
			h[i][j] = max(0, h[i-1][j], h[i][j-1], cost(p[i], t[j]))
		}

		if i+1 < len(p) {
			go populate(i+1, j)
		}
		if j+1 < len(t) {
			go populate(i, j+1)
		}
	}

	go populate(0, 0)
	fmt.Println(h[5][5])
}

// Compute the Smith-Waterman alignment matrix. Not optimized for parallelism
// because of the dependency of matrix formulation.
// @j6k4m8
func SW(p string, t string) alignment {

	// Create our matrix, called h by convention.
	h := make([][]int, len(p))
	for i := range h {
		h[i] = make([]int, len(t))
	}
	// Populate the matrix:
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(t); y++ {
			if x == 0 || y == 0 {
				h[x][y] = 0
			} else {
				h[x][y] = max(0, h[x-1][y], h[x][y-1], cost(p[x], t[y]))
			}
		}
	}

	// Compute the best reverse-path of an alignment matrix, starting with the
	// cell at max(x), max(y).
	y := len(t)
	x := len(p)

	path := ""

	// In the ambiguous case of more than one value being a possibility, this
	// algo 'tries' to minimize y as quickly as possible.
	for y > 0 {
		switch min(h[x-1][y], h[x-1][y-1], h[x][y-1]) {
		case h[x-1][y-1]:
			path = string(p[x]) + path
			x -= 1
			y -= 1
		case h[x][y-1]:
			y -= 1
			path = "-" + path
		case h[x-1][y]:
			path = string(p[x]) + path
			x -= 1
		}
	}

	// Now we store the sequence in path, and we store the offset in the
	// remaining value of y.
	return alignment{seq: path, offset: y}
}

func max(i, j, k, l int) int {
	if i > j && i > k && i > l {
		return i
	} else if j > k && j > l {
		return j
	} else if k > l {
		return k
	}
	return l
}

func min(i, j, k int) int {
	if i < j && i < k {
		return i
	} else if j < k {
		return j
	}
	return k
}
