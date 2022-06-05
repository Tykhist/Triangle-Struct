package main

import "fmt"

type Triangle struct {
	height float32
	base float32
}

func (triangle Triangle) area() float32{
 	return (triangle.base * triangle.height) / 2
}

func (triangle *Triangle) updateBase(newBase float32){
	triangle.base = newBase
}

func main() {
	triangle := Triangle{10, 4}
	
 	fmt.Println(triangle)
	fmt.Println(triangle.area())
	
	triangle.updateBase(13)
	fmt.Println(triangle)
	fmt.Println(triangle.area())
}
