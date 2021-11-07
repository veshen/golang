package main

import (
	"fmt"
	"os"
)

func main()  {
	fileObj ,err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file error",err)
	}
	// 关闭
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Println("close file error",err)
		}
	}(fileObj)

	tmp := make([]byte,128)
	for  {
		n, err := fileObj.Read(tmp[:])
		if err != nil{
			fmt.Println("read file error",err)
			return
		}
		// 读了多少字节
		fmt.Println(n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
            break
        }
	}
}
