package main

import (
	//"fmt"
	"os"
)


func main() {
	panic("a problem occured!")

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic("err")
	}
}
