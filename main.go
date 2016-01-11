package main

import (
	"./shapes"
	"fmt"
)

//Second Object Circle
type Circle struct {
	radius float32
}

func (s *Circle) Area() string {
	return fmt.Sprintf("%f", 3.14*s.radius*s.radius)

}

//Interface for shape, which creates a string
type Shape interface {
	Area() string
}

func printArea(s Shape) string {
	return fmt.Sprintf("area of square = %s \n", s.Area())
}

func main() {
	fmt.Println("hello world")

	testSquare := &shapes.Square{Side: 4}

	// testSquare := *shapes.Square{side: 5}
	testCircle := &Circle{radius: 5}
	// fmt.Println(testSquare)

	// fmt.Printf("area of square = %v \n", testSquare.area())

	fmt.Println(printArea(testSquare))
	fmt.Println(printArea(testCircle))

}
