package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data int
}

var instance *Singleton
var once sync.Once // 内核信号

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{ data: 10 }
	})
	return instance
}

func main() {
	i1 := GetInstance()
	i2 := GetInstance()
	i3 := &Singleton{ data: 10 }

	fmt.Printf("%x\n", &i1)
	fmt.Printf("%x\n", &i2)
	fmt.Printf("%x\n", &i3)
	fmt.Println(i1) // &{10}
	fmt.Println(i2) // &{10}
	fmt.Println(i3) // &{10}

	fmt.Println( i1 == i2) // true
	fmt.Println( i1 == i3) // false
	fmt.Println( i2 == i3) // false
}