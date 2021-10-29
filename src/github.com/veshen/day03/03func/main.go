package main

import "fmt"

// defer 延迟执行

func add(a int) func(b int) int {
	return func(b int) int{
		return a + b
	}
}

func main()  {

	fmt.Println(add(1)(2))
}
