package main

import "fmt"

func main()  {
	//a1 := []int { 1, 2, 3 }
	//a2:=make([]int,len(a1))
	//copy(a2, a1)
	//fmt.Println(a1,a2)

	x1 := [...]int { 1, 2, 3, 4, 5, 6, 7 }
	s1 := append(x1[2:3],x1[4:]...)
	fmt.Println(s1,x1)
}
