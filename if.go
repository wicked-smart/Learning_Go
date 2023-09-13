package main

import "fmt"

func main(){
	var a int;

	fmt.Scanln(&a)

	if a%2 == 0 {
		fmt.Println("Entered number is even..")
	} else {
		fmt.Println("Entered Number is Odd!!!")
	}

}

