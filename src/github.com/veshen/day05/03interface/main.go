package main

import "fmt"


// speaker type interface includes speak method
type speaker interface {
	speak()
}

// cat type struct includes name
type cat struct {
    name string
}

// dog type struct includes name
type dog struct {
    name string
}

func (c cat)speak()  {
	fmt.Println("miao miao miao")
}

func (d dog)speak()  {
	fmt.Println("wang wang wang")
}

func da(x speaker)  {
	x.speak()
}

func main()  {
	var c1 = cat{"tom"}
	var d1 = dog{name: "jerry"}
	da(c1)
	da(d1)
}
