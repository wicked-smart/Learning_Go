package main

import (
	"time"
	"fmt"
)


func main(){

	p := fmt.Println
	
	now := time.Now()

	p("now := ", now)

	then := time.Date(
		2023, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))



}

