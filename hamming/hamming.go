package hamming

import (
	"errors"
)

// Distance check dna length and compare differences between them
func Distance(a, b string) (int, error) {
	switch {
	case a == b:
		// a and b equal
		return 0, nil
	case len(a) != len(b):
		// cant compare differences dna
		return 0, errors.New("length of dns is different")
	default:
		// compare letter by letter and count differences
		diff := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		// return the hammig diferences
		return diff, nil
	}
}
