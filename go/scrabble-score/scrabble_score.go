package scrabble

import "unicode"

// Score calculates a word's value in the game scrabble
func Score(word string) int {
	score := 0

	for _, letter := range word {
		switch unicode.ToLower(letter) {
		case 'b':
			score += 3
		case 'c':
			score += 3
		case 'd':
			score += 2
		case 'f':
			score += 4
		case 'g':
			score += 2
		case 'h':
			score += 4
		case 'j':
			score += 8
		case 'k':
			score += 5
		case 'm':
			score += 3
		case 'p':
			score += 3
		case 'q':
			score += 10
		case 'v':
			score += 4
		case 'w':
			score += 4
		case 'x':
			score += 8
		case 'y':
			score += 4
		case 'z':
			score += 10
		default:
			score++
		}
	}
	return score
}
