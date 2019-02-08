package raindrops

import "strconv"

// Convert number to string
func Convert(num int) string {
	message := ""
	if num%3 == 0 {
		message += "Pling"
	}
	if num%5 == 0 {
		message += "Plang"
	}
	if num%7 == 0 {
		message += "Plong"
	}

	if message == "" {
		message = strconv.Itoa(num)
	}
	return message
}
