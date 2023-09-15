package main

import "fmt"

func zeroval(a int)  {
	a = 0
	//return a
}

// pass as reference
func oneval(a *int)  {
	*a = 1

	//return a
}


func main() {
	
	var a int = 34
	
	fmt.Println("a := ", a)

	zeroval(a);
	fmt.Println("zeroval a := ", a)

	oneval(&a);
	fmt.Println("oneval a := ", a)



}


