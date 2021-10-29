package main

import "fmt"

func sum(x int,y int)(ret int)  {
	return x + y
}

func main()  {
	fmt.Println(sum(1,2))
}
