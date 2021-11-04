package main

import "fmt"

// Dog type struct includes name and age
type Dog struct {
    name string
    age  int
}

// NewDog returns a new Dog
func NewDog(name string, age int) *Dog {
    return &Dog{name: name, age: age}
}


// Hwang is a Dog
// Dog 不带星号是值接收 是dog的拷贝
func (d Dog) Hwang() {
	fmt.Printf("%s www",d.name)
}

// Dog add age
// Dog带星号是接收的是指针
func (d *Dog) guonian() {
    d.age += 1
}

func main(){
	d1 := NewDog("xiaomi",1)
	d1.Hwang()
	d1.guonian()
	fmt.Println(d1.age)
}