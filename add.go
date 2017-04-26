package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main(){
    a := add(5, 11)
    fmt.Println(a)
}