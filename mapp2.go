package main

import "fmt"

func main(){

	foo := make(map[string]interface{})

	foo["name"] = "prem kumar"
	foo["age"] = 29
	foo["net worth"] = 2345.87


	fmt.Println("foo := ", foo)

	value, exists := foo["name"]
	
	fmt.Println("exists := ", exists)

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("not present")
	}

}






