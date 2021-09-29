package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	var s = "1"
	fmt.Printf("%T\n", n)  //查看类型
	fmt.Printf("%v\n", n)  //输出所有类型
	fmt.Printf("%#v\n", n) //输出所有类型 并且加描述
	fmt.Printf("%b\n", n)  //二进制
	fmt.Printf("%d\n", n)  //八进制
	fmt.Printf("%o\n", n)  //十进制
	fmt.Printf("%x\n", n)  //十六进制
	fmt.Printf("%s\n", s)  //字符串
}
