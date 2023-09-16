package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	
	s := "sha256 is the string!"

	h := sha256.New()

	h.Write([]byte(s))


	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs)

}
