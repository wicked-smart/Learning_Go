package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	
	//sort fruits by its length
	fruits := []string{"peach", "banana", "apple"}

	CmpByLen := func(a , b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, CmpByLen)

	fmt.Println("sorted fruits := ", fruits)

	//sort Person by age

	type Person struct {
		name string
		age int
	}
		
	people := []Person{
		Person{name: "Prem", age: 29},
		Person{name: "preeti", age: 27},
		Person{name: "Rinku", age: 47},
	}

	slices.SortFunc(people, func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)
}




		
