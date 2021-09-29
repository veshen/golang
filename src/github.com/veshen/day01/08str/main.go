package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串 必须用双引号包裹
	// s1 := "hello world"
	//字符 用单引号包裹
	// s2 := 'a'
	// s3 := '1'
	// s4 := '王'
	//字节 1字节=8Bit(8个二进制位) = 八个01010101
	//1个字符a = 1个字节
	//1个utf-8编码的汉字 一般占3个字节
	// s5 := `
	// 	多行字符串
	// 	第二行
	// 	第三行
	// `

	//字符串的常用操作
	// len(str) 求长度
	// + || fmt.Sprintf() 拼接字符串
	name := "veshen"
	age := "30"
	ss := fmt.Sprintf("%s%s\n", name, age)
	fmt.Printf("%s", ss) //veshen30

	//字符串分割
	str1 := "a/b/c"
	val := strings.Split(str1, "/")
	fmt.Println(val)
	fmt.Printf("%T\n", val)
	//判断包含
	fmt.Println(strings.Contains(str1, "a"))
	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(str1, "a"))
	fmt.Println(strings.HasSuffix(str1, "c"))
	//判断字符串中字符出现的位置
	fmt.Println(strings.Index(str1, "c"))
	//拼接 join
	fmt.Println(strings.Join(val, "---"))

}
