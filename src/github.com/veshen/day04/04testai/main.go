package main

import "fmt"

// Person type struct
type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "Veshen",
		Age:  18,
	}
	fmt.Println(p1)
}
