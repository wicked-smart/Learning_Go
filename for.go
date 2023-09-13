package main

import "fmt"

func main() {
    for i := 1; i <= 6; i++ {
        fmt.Println("i := ", i)
    }

    fmt.Println("printing star pattern ......")
  

    for i:=1; i <= 5; i++ {
        for j := 0; j < i; j++ {
            fmt.Print('*')
        }
        fmt.Println()
        i++
    }
}

