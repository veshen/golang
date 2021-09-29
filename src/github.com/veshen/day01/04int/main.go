package main

import "fmt"

func main() {
	var i1 = 101
	//十进制 %d  八进制 %o 二进制%b  十六进制 %x
	fmt.Printf("%d\n", i1)
	//八进制
	var i2 = 077
	fmt.Printf("%d\n", i2)
	//十六进制
	var i3 = 0xd7271
	fmt.Printf("%d\n", i3)

	//查看变量类型 %T
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	var i4 = int8(8)
	fmt.Printf("%T\n", i4)

	// int16 int32...

}
