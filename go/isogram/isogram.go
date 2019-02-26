package isogram

import "unicode"

// IsIsogram checks if word is isogram
func IsIsogram(word string) bool {
	count := make(map[rune]int)

	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		lower := unicode.ToLower(char)

		count[lower]++
		if count[lower] > 1 {
			return false
		}
	}
	return true
}
