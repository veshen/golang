package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件写内容
func wrDemo1()  {
	f,err := os.OpenFile("./xx.txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE,0644)
	if err != nil {
		fmt.Println("open file err")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("close file err")
		}
	}(f)
	fmt.Println("开始写入")
	f.Write([]byte("test"))
	f.WriteString("away")
	fmt.Println("写入完成")
}

func wrDemo2()  {
	f, err := os.OpenFile("./xx2.txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE,0644)
	if err != nil {
        fmt.Println("open file err")
    }
	defer func(f *os.File) {
        err := f.Close()
        if err != nil {
            fmt.Println("close file err")
        }
    }(f)

	wr := bufio.NewWriter(f)
	wr.WriteString("veshen")
	// 将缓存中的内容写入文件
	wr.Flush()
}

func main()  {
	//wrDemo1()
	wrDemo2()
}
