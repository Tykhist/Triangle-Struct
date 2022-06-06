package main

import "fmt"

type Triangle struct {
	// Creates a data type with fields for a triangle
	height float32
	base float32
}

func (triangle Triangle) area() float32 {
	// Calculates the area of the triangle
 	return (triangle.base * triangle.height) / 2
}

func (triangle *Triangle) updateBase(newBase float32) {
	// Updates the base field using a pointer for the Triangle struct
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
