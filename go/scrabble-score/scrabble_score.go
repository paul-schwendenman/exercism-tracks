package scrabble

import "unicode"

// Score calculates a word's value in the game scrabble
func Score(word string) int {
	score := 0

	for _, letter := range word {
		switch unicode.ToLower(letter) {
		case 'q', 'z':
			score += 10
		case 'j', 'x':
			score += 8
		case 'k':
			score += 5
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'd', 'g':
			score += 2
		default:
			score++
		}
	}
	return score
}
