package cmd

import (
	"encoding/json"
	"fmt"
)


// An alignment stores information about the actual string as well
// as its position.
type alignment struct {
	offset int,
	seq string
}


// Calculate the "cost" (for insertion into alignment matrices) of two strings
// @j6k4m8
func cost(c0 string, c1 string) int {
    // Using the scores described by Smith et al (1981). This is effectively
	// changed to { 0, 2 } in the SW implementation because we max against 0.
    if c0 == c1 {
        return 2;
    } else {
        return -1;
    }
}

// Compute the Smith-Waterman optimal alignment
func AlignSW(p string, t string) int {

    // Create our matrix, called h by convention.
    h := [len(p)][len(t)]int{}

    for x := 0; x < len(p); x++ {
        for y := 0; y < len(t); y++ {
            if x == 0 || y == 0 {
                h[x][y] = 0;
            } else {
                h[x][y] = max(0, h[x-1][y], h[x][y-1], cost(p[x], t[y]))
				fmt.Printf("%d\n", h[x][y])
            }
        }
    }

}
