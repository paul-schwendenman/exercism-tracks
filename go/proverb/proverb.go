package proverb

import "fmt"

// Proverb generator
func Proverb(rhyme []string) (poem []string) {
	rhymeLen := len(rhyme)

	if rhymeLen == 0 {
		return
	}

	poem = make([]string, rhymeLen)

	for pos := 0; pos < rhymeLen-1; pos++ {
		poem[pos] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[pos], rhyme[pos+1])
	}

	poem[rhymeLen-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return
}
