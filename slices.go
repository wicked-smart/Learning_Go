package main

import "fmt"

func main(){

	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)
	
	s = make([]string, 3)
	fmt.Println("emp:= ", s, "len: ", len(s), "cap: ", cap(s))
	
	s[0] = "who is this guy"
	s[1] = "i need to go to Malhad!!!"
	
	s = append(s, "hoooha!!!")
	fmt.Println("s := ", s)

	fmt.Println("s[1]:= ", s[1])

	
	var c []string
	c = make([]string, len(s))

	copy(c,s)

	l := c[0:2]

	fmt.Println(l)

	
}
