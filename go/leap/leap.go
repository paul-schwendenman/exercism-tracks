package leap

// IsLeapYear returns true if the year would have an extra day
func IsLeapYear(year int) bool {
	return isDivisbleBy(year, 4) && (!isDivisbleBy(year, 100) || isDivisbleBy(year, 400))
}

func isDivisbleBy(number int, divisor int) bool {
	return number%divisor == 0
}
