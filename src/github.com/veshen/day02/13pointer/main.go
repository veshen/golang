package main

import "fmt"

// pointer
// &: 取地址
// *: 根据地址取值

func main()  {

	n := 123
	fmt.Println(&n)

	p := *&n
	fmt.Println(p)

}
