package sort

import (
	"fmt"
	"sort"
)

func Run() {
	fmt.Println("Running sort program")

	// declare input slice
	inputSlice := []int{1, 8, 0, 6, 2, 9, 5, 3, 7, 4}
	fmt.Println("Input int slice: ", inputSlice)

	fmt.Println("Sorting the input int slice")
	sort.Ints(inputSlice)
	fmt.Println("Sorted int slice: ", inputSlice)

	// check if the int slice is sorted
	fmt.Println("Is input int slice sorted: ", sort.IntsAreSorted(inputSlice))

	x := 12
	// search x in the inputSlice
	id := sort.SearchInts(inputSlice, x)
	fmt.Printf("Index of %d is %d\n", x, id)

	i := sort.Search(len(inputSlice), func(i int) bool { return inputSlice[i] >= x })
	fmt.Printf("Index of %d is %d\n", x, i)

	// declare input string
	inputString := []string{"harsh", "zebra", "sameer", "alice", "bob", "marley"}
	fmt.Println("Input string slice: ", inputString)

	fmt.Println("Sorting the input string slice")
	sort.Strings(inputString)
	fmt.Println("Sorted string slice: ", inputString)

	// check if the string slice is sorted
	fmt.Println("Is input string slice sorted: ", sort.StringsAreSorted(inputString))
}
