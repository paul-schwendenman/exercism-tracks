package proverb

import "fmt"

// Proverb generator
func Proverb(rhyme []string) (poem []string) {
	if len(rhyme) == 0 {
		return
	}

	for pos := 0; pos < len(rhyme)-1; pos++ {
		poem = append(poem, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[pos], rhyme[pos+1]))
	}

	poem = append(poem, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return
}
