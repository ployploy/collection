package fizzbuzz

import (
	"strconv"
)

func Say(number int) string {
	maps := make(map[int]string)
	maps[3] = "Fizz"
	maps[5] = "Buzz"
	maps[15] = "FizzBuzz"
	arrays := []int{15, 3, 5}
	for _, element := range arrays {
		if number%element == 0 {
			return maps[element]
		}
	}

	return strconv.Itoa(number)
}
