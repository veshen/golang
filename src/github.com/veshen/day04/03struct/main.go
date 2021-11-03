package main

import "fmt"

type person struct {
	name string
	age int
	hobby []string
}



func main()  {
	var veshen person
	veshen.name = "wang"
	veshen.age = 30
	veshen.hobby = []string {"aaa"}

	fmt.Println(veshen)
	fmt.Println(veshen.name)

	var man = person{
		name: "aaa",
		age: 123,
	}
	fmt.Println(man)

}
