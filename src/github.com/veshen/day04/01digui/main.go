package main

import "fmt"

func digui(n uint) uint{
	if n<=1{
		return 1
	}
	return n * digui(n-1)
}

func main()  {
	num := digui(10)
	fmt.Println(num)
}




