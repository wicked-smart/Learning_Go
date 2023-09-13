package main

import "fmt"

func main(){

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8
	m["k3"] = 9

	fmt.Println("map := ", m)

	fmt.Println("m[k2] := ", m["k2"])

	// delete m["k2"]

	delete(m, "k2")
	fmt.Println("map after deleting k2 := ", m)

	
	foo, prs := m["k3"]

	fmt.Println("foo  prs := ",foo,  prs)

}
