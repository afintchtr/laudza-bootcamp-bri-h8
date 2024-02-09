package funcs

import "fmt"

type callbackOdd func(int) bool

func findOddNumbers(numbers []int, callback callbackOdd) int {
	totalOddNumbers := 0
	for _, value := range numbers {
		if callback(value) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

func DoThat(randomNums func() []int) {
	var find = findOddNumbers(randomNums(), func(i int) bool {
		return i%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}