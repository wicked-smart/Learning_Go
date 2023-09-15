package main

import "fmt"

func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}


func main() {
	
	sum(1,2)
	sum(1,2,3)

	nums := []int{23,21,34,4,5}
	
	sum(nums...)
}


