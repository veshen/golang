package main

import "fmt"

func main()  {
	// 元素类型为map的切片
	var m1 = make([]map[int]string,10,100) // 定义 初始化切片
	m1[0] = make(map[int]string,1)
	m1[0][1] = "v"

	fmt.Println(m1)

}
