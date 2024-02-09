package main

import (
	"fmt"
	"golang-days/day06/funcs"
	"math"
	"math/rand"
	"strconv"
)

func whichIsClosestToTen(num1, num2 int) int {
	diff1 := math.Abs(float64(10.0 - num1))
	diff2 := math.Abs(float64(10.0 - num2))
	if diff1 < diff2 {
		return num1
	} else if diff2 < diff1 {
		return num2
	}
	return 0
}
func displayWhichIsClosestToTen() {
	paramA := rand.Intn(100 - (-100)) + (-100)
	paramB := rand.Intn(100 - (-100)) + (-100)
	result := whichIsClosestToTen(paramA, paramB)
	fmt.Printf("closest number to 10 between %v and %v is %v", paramA, paramB, result)
}

func whichIsClosestToTenVariadic(nums ...int) int {
	isEveryDiffIsSame := true
	closest := math.MaxFloat64
	var result int
	for index, num := range nums {
		diff := math.Abs(float64(10.0 - num))
		if diff < float64(closest) {
			closest = diff
			result = num
		}
		if index > 0 && diff != closest {
			isEveryDiffIsSame = false
		}
	}
	if isEveryDiffIsSame {
		return 0
	}
	return result
}

func displayWhichIsClosestToTenVariadic() {
	result := whichIsClosestToTenVariadic(9, 11, 9, 11, 9, 9, 9)
	if result == 0 {
		fmt.Println("every diffs are equal!")
	} else {
		fmt.Println("the closest num to ten is", result)
	}
}

func findThreeConsecutive(nums ...int) bool {
	for index, num := range nums {
		if index < 2 {
			continue
		}
		currNum := &num
		if *currNum == nums[index - 1] && *currNum == nums[index - 2] {
			return true
		}
	}
	return false
}

func findThreeConsecutiveInClosure() {
	findThreeConsecutive := func (nums ...int) bool {
		for index, num := range nums {
			if index < 2 {
				continue
			}
			currNum := &num
			if *currNum == nums[index - 1] && *currNum == nums[index - 2] {
				return true
			}
		}
		return false
	}
	fmt.Println(findThreeConsecutive(9, 9, 9, 11, 9, 11, 9, 9, -9))
}

func extractLastDigit(nums ...int) func() []int{
	someFunction := func() []int {
		var result []int

		for _, num := range nums {
			stringNum := strconv.Itoa(num)
			lastDigit := stringNum[len(stringNum)-1:]
			intLastDigit, _ := strconv.Atoi(lastDigit)
			if intLastDigit != 0 {
				result = append(result, intLastDigit)
			}
		}
		return result
	}
	return someFunction
}

func nestedDoThat(n int) {
	funcs.DoThat(func() []int {
		var nums []int
		for i := 0; i < n; i++ {
			nums = append(nums, rand.Intn(200))
		}
		fmt.Println(nums)
		return nums
	}) 
}

func main() {
	funcs.PracticeReflectInStruct()
}
