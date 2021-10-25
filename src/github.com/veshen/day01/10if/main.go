package main

import "fmt"

func main() {
	//age := 19
	//if age > 10 {
	//	fmt.Print(age)
	//}else{
	//	fmt.Println("小")
	//}
	//
	//if name := "veshen"; name == "veshen" {
	//	fmt.Println(name)
	//}

	//for循环
	//for a := 0; a < 10; a ++ {
	//	fmt.Println(a)
	//}
	//
	//a2 := 0
	//for ; a2 < 10; a2 ++ {
	//	fmt.Println(a2)
	//}
	//
	//a3 := 0
	//for ; a3 < 10; {
	//	a3++
	//}

	//for range 循环
	//name := "veshen"
	//for i,v := range name {
	//	fmt.Printf("%d %c\n",i,v)
	//}

	for start := 1; start < 10; start++ {
		for end := 1; end <= start; end++ {
			if start == end {
				fmt.Printf("%d * %d = %d\n", start, end, start*end)
			} else {
				fmt.Printf("%d * %d = %d ", start, end, start*end)
			}
		}
	}
}
