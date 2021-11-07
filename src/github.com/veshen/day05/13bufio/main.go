package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// read file use bufio
func readFile()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
        fmt.Println("open file error")
    }
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Println("close file error")
		}
	}(fileObj)

	// read file use bufio
	reader := bufio.NewReader(fileObj)
	var fileContent string
	for {
        str, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
		fileContent += str
    }
	fmt.Print(fileContent)


}

// read file by ioutil
func readFileByIoutil()  {
	ret, err := ioutil.ReadFile("./main.go")

	if err != nil {
        fmt.Println("read file error")
    }
	fmt.Print(string(ret))
}

func main()  {
	//readFile()
	readFileByIoutil()
}
