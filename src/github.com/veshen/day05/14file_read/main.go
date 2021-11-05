package main

import (
	"fmt"
	"os"
)

// 打开文件写内容


func main()  {
	f,err := os.OpenFile("./xx.txt",os.O_APPEND|os.O_CREATE,0666)
	if err != nil {
        fmt.Println("open file err")
    }
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("close file err")
		}
	}(f)
	f.Write([]byte("test"))
	f.WriteString("away")

}
