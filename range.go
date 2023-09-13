package main

import "fmt"

func main(){

	nums := [5]int{2,3,4,5,34}
	
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum := ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index := ", i)
		}
	}


	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k,v := range kvs {
		fmt.Println("key := %s -> value := %s", k, v)
	}

	str := "testing strings"

	for idx, c := range str {
		fmt.Println(" str['%s'] = %s", idx, c)
	}

	//only get the keys of kvs
	for value := range kvs {
		fmt.Println("%s", value)
	}
}
	
