package main

import (
	"fmt"
	"unsafe"
)

type s struct {
	b [10]int
	a [3]bool
	c []int
	k chan int
	j uintptr
}

func main() {
	a := s{a: [...]bool{false, false, false}, b: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, c: []int{1, 2}}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Alignof(a.a))
	fmt.Println(unsafe.Alignof(a.b))
	fmt.Println(unsafe.Alignof(a.c))
	fmt.Println(unsafe.Alignof(a.k))
	fmt.Println(unsafe.Alignof(a.j))
}
