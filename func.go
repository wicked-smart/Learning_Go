package main

import "fmt"

func add(a int, b int) int {
	return a + b
}


func addadd(a, b, c int) int {
 	
	return a + b + c
}


func main() {
	
	fmt.Println("sum := ", add(23,45))
	fmt.Println("summm := ", addadd(12,23,34))

}
