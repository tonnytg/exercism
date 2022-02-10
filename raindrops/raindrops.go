package raindrops

import "strconv"

// Convert converts a number to a string
func Convert(number int) string {

	// msg is the string that will be returned
	msg := ""

	// if the number is divisible by 3, add "Pling" to the msg
	if number % 3 == 0 {
		msg += "Pling"
	}

	// if the number is divisible by 5, add "Plang" to the msg
	if number%5 == 0 {
		msg += "Plang"
	}

	// if the number is divisible by 7, add "Plong" to the msg
	if number%7 == 0 {
		msg += "Plong"
	}

	// if the number is not divisible by 3, 5, or 7, return the number as a string
	if msg == "" {
		msg = strconv.Itoa(number)
	}
	return msg
}
