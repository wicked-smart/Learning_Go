package main

import "fmt"


type Person struct {
	name string
	age int
}


func newPerson(name string) *Person {
	
	p := Person{name: name}
	p.age  = 29

	return &p
}


func main() {
	
	fmt.Println(Person{"Prem Kumar", 29})
	fmt.Println(Person{name: "Alice", age: 30 })
	fmt.Println(Person{name: "Fred"})
	fmt.Println(newPerson("Jon"))
}
