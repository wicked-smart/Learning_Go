package main

import (
	"fmt"
	b64 "encoding/base64"
)

func main(){
	
	data := "prem kumar will not give up  "

	base64 := b64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("base64 : ", base64)

	decode, _ := b64.StdEncoding.DecodeString(base64)

	fmt.Println(string(decode))
	fmt.Println()


}


