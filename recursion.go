package main

import "fmt"


var rec = make(map[int]int)

func fib(a int) int {
	
	if a == 0 || a == 1 {
		return 1
	} 

	value, exists := rec[a]

	if exists {
		return value
	} else {
		rec[a] = fib(a-1) + fib(a-2)
		return rec[a]
	}

}



func main(){
	
	// find fibonaci of a given number
	var num int

	fmt.Scanln(&num)

	fmt.Println(fib(num))

	fmt.Println("recursion map := ", rec)

}
