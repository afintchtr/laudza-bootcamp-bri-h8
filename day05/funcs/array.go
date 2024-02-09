package funcs

import "fmt"

func AnArray() {
	myName := "afin"

	var arr[4][5]string = [4][5]string{
		{"alvin", "arif", "reza", "rinaldi", "nina"}, 			// 0
		{"noel", "dilla", "rosa", "juan michel", "teguh"},	// 1
		{"septyan", "farras", "giyanda", "afin", "azwar"},	// 2
		{"dionysius", "rifki", "raffi", "zain"},						// 3
	}

	for firstIndex, firstValue := range arr {
		for secondIndex, secondValue := range firstValue {
			if (secondValue == myName) {
				fmt.Println(firstIndex, firstValue, secondIndex, myName)
			}
		}
	}

	var fruits = []string{"apple", "orange", "banana"}
	var fruits2 = fruits[1:]
	fmt.Printf("%#v\n", fruits2)
}
