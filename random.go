package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64()*5 + 5)

	//non-deterministic changes using seeding
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))
	

}
