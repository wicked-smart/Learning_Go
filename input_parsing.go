package main

import (
	"fmt"
	"strconv"
)

func main(){
	
	f, test := strconv.ParseFloat("1.234", 64)

	p := fmt.Println


	p(f)
	p(test)

	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)

	h, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(h)

	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)

	k, _ := strconv.Atoi("135")
	p(k)

	foo, err := strconv.Atoi("wat")
	p(foo)
	p(err)

}
