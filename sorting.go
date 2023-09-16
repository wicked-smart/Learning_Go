package main

import (
	"fmt"
	"slices"
)

func main() {
	
	strs := []string{"c", "a", "b", "d"}
	slices.Sort(strs)

	fmt.Println("strs := ", strs)

	ints := []int{21,3,6,87}
	fmt.Println("Before sorting IsSorted := ", slices.IsSorted(ints))
	slices.Sort(ints)

	fmt.Println("Int Sorted := ", ints)

}


