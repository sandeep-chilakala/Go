package main

import "fmt"

//struct
type demoStruct struct {
	name string
	age  int
}

//function with pointer as input
func update(str *demoStruct) {
	str.name = "sandeep"
	str.age = 28
}

//method of type demostruct
func (d *demoStruct) update(str string, i int) {
	d.age = i
	d.name = str
}

func main() {
	var x []int = []int{1, 2, 3}
	y := make([]int, len(x))
	copy(y, x)
	y[0] = 10
	fmt.Println(x, y)

	var p demoStruct
	p.update("hello", 30)
	fmt.Println(p)
	update(&p)
	fmt.Println(p)
}
