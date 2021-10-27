package main

import "fmt"

// 切片 slice

func main()  {
	var s = []int { 1, 2, 3 } // 定义一个int类型的切片
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// 由数组得到切片
	var a1 = [...]int { 1,2,3 }
	a2 := a1[0:2] //左闭右开
	fmt.Println(a2) // [ 1, 2 ]
}
