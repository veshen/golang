package main

import "fmt"

func main() {
	var n = 1

	switch n {
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("none")
	}

	switch a:=1;a {
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")
	}
	
	a := 30

	switch  {
	case a>10:
		fmt.Println("a>10")
	default:
		fmt.Println("d")
		
	}

}
