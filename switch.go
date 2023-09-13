package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println(time.Now().Weekday())

	start := time.Now()

	time.Sleep(3 * time.Second)

	t := time.Now()
	
	elapsed := t.Sub(start)

	fmt.Println("Total Elapsed time is ", elapsed)

}


