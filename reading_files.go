package main

import (
	"fmt"
	//"bufio"
//	"io"
	"os"
)

func main() {
	
		dat, err := os.ReadFile("foo.txt")

		if err != nil {
			panic(err)
		}

		fmt.Println(string(dat))

		// create a file and 
		f, err := os.Open("foo.txt")

		if err != nil {
			panic(err)
		}

		b1 := make([]byte, 5)

		n1, err := f.Read(b1)

		if err != nil {
			panic(err)
		}

		fmt.Println("%d bytes: %s\n", n1, string(b1[:n1]))
		
		
		//placing file reader at a place and seeking upto a place

		o2, err := f.Seeek(6, 0)
		if err != nil {
			panic(err)
		}

		b2 := make([]byte, 2)
		n2, err := f.Read(b2)

		if err != nil {
			panic(err)
		}

		fmt.Println("")

}


	
