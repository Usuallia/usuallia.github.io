package main

import (
	"fmt"
)

func iter(k int) func() (int, bool) {
	c := make(chan int)
	//done := false
	go func() {
		for i := 0; i < k; i++ {
			c <- i
			//fmt.Println(i)
		}
		//done = true
		close(c) // 一行close代替done标志的效果
	}()
	return func() (int, bool) {
		res, ok := <-c
		return res, ok
		//var res int
		//if !done {
		//	res = <-c
		//} else {
		//	res = 0
		//}
		//return res, done
	}
}

// <-chan只读通道,chan<-只写通道
func iterGrace(k int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < k; i++ {
			c <- i
		}
		close(c)
		// c <- 2 向close的通道发送数据会引起panic
		// 应该了解panic和recover的机制（类似于exception与try catch机制）
	}()
	return c
}
func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("panic as follow:", r)
	//	}
	//}()
	defer fmt.Println("Delay done")

	it := iter(5)
	for true {
		if i, ok := it(); ok {
			fmt.Println(i)
		} else {
			break
		}

	}

	grace := iterGrace(5)
	for i := range grace { // range自动判别channel的close状态
		fmt.Println(i)
	}

	//time.Sleep(time.Second)
	panic("wtf?")
	fmt.Println("Did I live?")
	fmt.Println("I suppose no")
}
