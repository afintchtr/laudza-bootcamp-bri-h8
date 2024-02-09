package funcs

import (
	"fmt"
)

func main() {
	thisIsSlice := []int{1, 2};
	fmt.Println(thisIsSlice, len(thisIsSlice), cap(thisIsSlice))
	fmt.Printf("%p\n", thisIsSlice)
	fmt.Printf("%p\n", &thisIsSlice)
	
	thisIsSlice = append(thisIsSlice, 3)
	fmt.Println(thisIsSlice, len(thisIsSlice), cap(thisIsSlice))
	fmt.Printf("%p\n", thisIsSlice)
	fmt.Printf("%p\n", &thisIsSlice)

	// fmt.Printf()
}