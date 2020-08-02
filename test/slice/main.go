package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3, 4)
	printSlice(a, "before")
	changeSlice(a, 1)
	printSlice(a, "after")

}



func changeSlice(s []int, n int) {
	s = append(s, n)
	printSlice(s, "changing")
}

func printSlice(s []int, position string) {
	fmt.Printf("[position:%-8s] [len:%d] [cap:%d] [pointer:%p] value:%v\n", position, len(s), cap(s), s, s)
}
