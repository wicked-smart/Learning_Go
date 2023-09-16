package main

import (
	"fmt"
	s "strings"
)

var p  = fmt.Println

func main(){
	
	p("Contains:  ", s.Contains("testing", "es"))
	p("count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("Index : ", s.Index("prem is my name", "i"))
	p("Join :  ", s.Join([]string{"p", "r", "e", "m"},"--" ))
	p("Replace :", s.Replace("prem is a good girl!", "prem", "preeti", 1))
	p("Split : ", s.Split("p-r-e-m", "-"))
	p("ToLower : ", s.ToLower("TEST"))
	p("ToUpper : ", s.ToUpper("TestingggG!"))

}


