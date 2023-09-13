package main

import "fmt"

func main(){

	var a [5]int
	
	a[4] = 90
	a[1] = 67

	for i:=0; i<5; i++ {
		fmt.Println(a[i])
	}

	fmt.Println("length of a := ", len(a))
	
	var twoD [4][5]int

	for i:=0; i<4; i++ {
		for j:=0; j<5; j++ {
			twoD[i][j] = i+j
		}
	}

	fmt.Println("2d:= ", twoD)

}
