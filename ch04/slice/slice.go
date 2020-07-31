package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:3]
	fmt.Printf("a:%p\nb:%p\n", a, b)
	fmt.Printf("a:%v\nb:%v\n", a, b)
	a[1] = 4
	fmt.Printf("a:%p\nb:%p\n", a, b)
	fmt.Printf("a:%v\nb:%v\n", a, b)
}
