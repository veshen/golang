package main

import "fmt"

func fn1()  {
	fmt.Println('a')
}

func fn2()  {
	defer func(){
		var err = recover()
		fmt.Println("error:",err)
	}()
	fmt.Println('b')
	panic("error") //程序崩溃退出
}

func fn3()  {
	fmt.Println('c')
}

func main()  {
	fn1()
	fn2()
	fn3()
}
