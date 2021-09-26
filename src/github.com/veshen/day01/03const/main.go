package main

import "fmt"

//常量
const (
	name string = "veshen"
	age  int    = 30
	isOk bool   = true
)

//批量声明常量时 如果n2没赋值 则n2 和上一行一样
const (
	n1 = 100
	n2 //100
	n3 //100
)

//iota
const (
	b1 = iota //0
	b2        //1
	b3        //2
)

const (
	a1 = iota //0
	a2        //1
	_         //2
	a3        //3
)

//const 每新增一行使iota+1
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

//定义数量级 1左移N位
const (
	_  = iota
	KB = 1 << (1 * iota)
	MB = 1 << (1 * iota)
	GB = 1 << (1 * iota)
	TB = 1 << (1 * iota)
	PB = 1 << (1 * iota)
)

func main() {
	fmt.Print(name, age, isOk)
}
