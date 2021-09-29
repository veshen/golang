package main

import "fmt"

//浮点数
func main() {
	// math.MaxFloat32

	//go语言中浮点数默认float64
	f1 := 123.123123

	//强制定义为float32
	f2 := float32(1.23)

	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
}
