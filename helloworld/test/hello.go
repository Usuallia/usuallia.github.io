package main

import (
	"fmt"
)

type People interface {
	echo()
}

type Student struct {
	id    string
	name  string
	age   int
	store float32
}

func (stu Student) echo() {
	fmt.Printf("My name is %s, my id is %s, i'm %d year's old and I have %f store\n", stu.name, stu.id, stu.age, stu.store)
}

func createMap() map[string]People {
	stu := &Student{
		id:    "3119",
		name:  "zk",
		age:   19,
		store: 3000.00,
	}
	gdu := Student{
		id:    "3120",
		name:  "zk",
		age:   19,
		store: 3000.00,
	}
	// 奇怪的是stu指针和stu都可以被转为接口类型
	gm := map[string]People{
		"3119": stu,
	}
	gm["3120"] = gdu
	return gm
}
func main() {
	//stu := new(Student)

	gm := createMap()
	for k, v := range gm {
		fmt.Println(k, v)
	}
	a, ok := gm["22"]
	fmt.Println(a, ok, a == nil)
	ints := make(chan int, 3)
	ints <- 3
	ints <- 2
	ints <- 1
	for len(ints) != 0 {
		fmt.Println(<-ints)

	}
}
