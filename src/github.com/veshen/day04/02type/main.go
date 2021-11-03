package main

import "fmt"

// 自定义类型和类型别名

type myInt int

func main()  {
	var n myInt = 100
	fmt.Println(n)
	fmt.Printf("%T",n)
}
