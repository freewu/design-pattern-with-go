package main

import "fmt"

type ProductInterface interface {
	Show() string
}

type Product1 struct {}
func (*Product1) Show() string {
	return "product-1"
}

type Product2 struct {}
func (*Product2) Show() string {
	return "product-2"
}

func ProductSimpleFactory(t string) ProductInterface {
	switch t {
	case "1":
		return &Product1{}
	case "2":
		return &Product2{}
	}
	return nil
}

func main() {
	f1 := ProductSimpleFactory("1")
	fmt.Println(f1)
	fmt.Println(f1.Show())
	f2 := ProductSimpleFactory("2")
	fmt.Println(f2)
	fmt.Println(f2.Show())
	f3 := ProductSimpleFactory("3")
	fmt.Println(f3) // nil
}