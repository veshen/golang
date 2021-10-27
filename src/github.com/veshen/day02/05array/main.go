package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放元素的类型和容量(长度)
// 数组长度是数组类型的一部分
func main(){
	//var a [3]bool //如果不初始化默认元素都是零值 布尔值就是false 整型 浮点型是0 字符串是""
	//a[0] = true
	//
	//a = [3]bool{true,true,true}
	//
	//var a2 = [...]int {1,2,3} // ... 数组长度用初始化的值推算
	//
	//a3 := [5]int {0:1,4:2} // [1,0,0,0,2] 根据索引初始化
	//
	//fmt.Println(len(a))
	//fmt.Printf("%T \n",a[0])
	//fmt.Printf("%T \n",a2)
	//fmt.Printf("%T \n",a3)

	//数组的遍历
	citys := [...]string { "北京", "上海", "深圳" }
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for _,v := range citys{
		fmt.Println(v)
	}
	// 多维数组
	// [ [1,2], [3,4], [5,6] ]
	var a5 = [3][2]int { {1,2}, {3,4}, {5,6} }
	fmt.Println(a5)

	//var a6 = [...]int {1,2,3,4,5};

}
