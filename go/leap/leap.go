package leap

// IsLeapYear returns true if the year would have an extra day
func IsLeapYear(year int) bool {
	return year%4 == 0 && (!(year%100 == 0) || (year%400 == 0))
}
