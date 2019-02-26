package isogram

import "unicode"

// IsIsogram checks if word is isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		lower := unicode.ToLower(char)

		if seen[lower] {
			return false
		}
		seen[lower] = true
	}
	return true
}
