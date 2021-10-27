package main

import "fmt"

func main()  {
	var s1 = []string { "北京", "上海" }
	s1 = append(s1, "深圳")
	fmt.Println(s1)
}
