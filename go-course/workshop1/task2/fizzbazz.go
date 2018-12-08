package task2

import (
	"strconv"
	"strings"
)

func FizzBuzz(count int) string {
	var o = ""

	for i := 1; i <= count; i++ {
		if (i%3 == 0) {
			o += "Fizz "
		}

		if (i%5 == 0) {
			o += "Buzz "
		}

		if (i%3 != 0 && i%5 != 0) {
			o += strconv.Itoa(i) + " "
		}
	}

	return strings.Trim(o, " ")
}