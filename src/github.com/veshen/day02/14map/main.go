package main

import "fmt"

func main()  {
	var m1 map[string]int //还没有在内存中开辟空间
	m1 = make(map[string]int,10) //在设计的时候估算好map的容量， 避免在运行中扩容
	m1["veshen"] = 30
	m1["other"] = 99

	fmt.Println(m1)

	value, ok := m1["vesh"]
	if ok {
		fmt.Println(value)
	}else{
		fmt.Println("nil")
	}
	
	// map的遍历
	for k, v := range m1 {
		fmt.Println(k,v)
	}

	//删除key delete(map,key)
	delete(m1, "other")

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k,v)
	}

}
