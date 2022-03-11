package proverb

import "fmt"

// Proverb return a empty slice of string or many phrases.
func Proverb(rhyme []string) (a []string) {

	// if slice of string had more than 1 word
	for i := 1; i < len(rhyme); i++ {
		// rhyme[i-1] -> get penultimate word
		// rhyme[i] -> get last word
		a = append(a, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	// get last word of slice
	if len(rhyme) > 0 {
		a = append(a, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	// return all
	return a
}
