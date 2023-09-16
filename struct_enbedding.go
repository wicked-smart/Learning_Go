package main

import (
	"fmt"
)


type base struct {
	num int
}

type container struct {
	base 
	str string
}

func main(){
	
	foo := container{
		base: base{
			num: 1,
		}, 
		str: "prem kumar is rocking htis worls !",
	}

	fmt.Println(foo)
}
