package main

import (
	"fmt"
	"helloworld/algorithm"
)

func Protect(runnable func()) bool {
	res := true
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurs")
			res = false
		}
	}()
	runnable()
	fmt.Println("still working")
	return res
}
func main() {
	fmt.Println(Protect(func() {
		panic("hei")
	}))
	fmt.Println(Protect(func() {
	}))
	algorithm.Solution()
}
