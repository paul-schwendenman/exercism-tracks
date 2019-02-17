package strain

// Ints is list of integers
type Ints []int

// Lists is a list of lists of ints
type Lists [][]int

// Strings is a list of strings
type Strings []string

// Keep filters through array removing ints that don't match condition
func (in Ints) Keep(operation func(int) bool) (out Ints) {
	for _, item := range in {
		if operation(item) {
			out = append(out, item)
		}
	}

	return
}

// Discard filters through array removing ints that match condition
func (in Ints) Discard(operation func(int) bool) (out Ints) {
	for _, item := range in {
		if !operation(item) {
			out = append(out, item)
		}
	}

	return
}

// Keep filters through array removing lists that don't match condition
func (in Lists) Keep(operation func([]int) bool) (out Lists) {
	for _, item := range in {
		if operation(item) {
			out = append(out, item)
		}
	}

	return
}

// Keep filters through array removing strings that don't match condition
func (in Strings) Keep(operation func(string) bool) (out Strings) {
	for _, item := range in {
		if operation(item) {
			out = append(out, item)
		}
	}

	return
}
