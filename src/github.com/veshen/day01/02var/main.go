package main

import "fmt"

//声明变量
var name string
var age int
var isOk bool

//批量生命变量

var (
	name2 string
	age2  int
	isOk2 bool
)

//匿名变量 ——

func foo() (string, int) {
	return "veshen", 30
}

func main() {

	name = "veshen"
	age = 30
	isOk = true
	//%s:占位符
	fmt.Print(isOk)
	fmt.Printf("name:%s", name)
	fmt.Println(age) //Println 自动加换行符


	//简短变量声明 := 只能在函数里面用
	s3 := "hahha"
	fmt.Print(s3)

	//匿名变量的使用
	nameX, _ := foo()
	_, ageY := foo()
	fmt.Println(nameX)
	fmt.Println(ageY)

}
