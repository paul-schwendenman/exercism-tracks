package accumulate

// Accumulate runs operation on list of strings
func Accumulate(target []string, operation func(string) string) []string {
	output := make([]string, len(target))
	for pos, word := range target {
		output[pos] = operation(word)
	}

	return output
}
