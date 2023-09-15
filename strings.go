package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	
	const s = "foobar"

	//print hexadecimal values of strings
	for idx, value := range s {
		fmt.Println(" s[%d] = %x ",idx, value )
	}
	
	//find rune count
	fmt.Println("Rune count := %d", utf8.RuneCountInString(s))

	//Decode Rune In String


}	
