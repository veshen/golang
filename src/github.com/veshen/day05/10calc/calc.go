package calc

import "fmt"

func init()  {
	fmt.Println("import 我时自动执行")
}

func Add(x, y int) int {
	return x + y
}
