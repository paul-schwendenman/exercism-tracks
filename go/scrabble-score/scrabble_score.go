package scrabble

import "strings"

var letterValue = map[rune]int{
	'b': 2,
	'c': 2,
	'd': 1,
	'f': 3,
	'g': 1,
	'h': 3,
	'j': 7,
	'k': 4,
	'm': 2,
	'p': 2,
	'q': 9,
	'v': 3,
	'w': 3,
	'x': 7,
	'y': 3,
	'z': 9,
}

// Score calculates a word's value in the game scrabble
func Score(word string) int {
	score := 0

	for _, letter := range strings.ToLower(word) {
		score += letterValue[letter] + 1
	}
	return score
}
