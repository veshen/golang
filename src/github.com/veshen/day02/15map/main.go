package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano());

	var sourceMap = make(map[string]int,200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Intn(100)
		sourceMap[key] = value
	}

	//fmt.Println(sourceMap)

	// 取出map中所有的key存入切片keys
	keys := make([]string,0,200);
	for key := range sourceMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, sourceMap[key])
	}
}
