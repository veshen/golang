package main

import "fmt"

func main(){
	//跳出for循环
	//break 结束for循环
	//continue 结束本次for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
	}
	fmt.Println("end")
}
