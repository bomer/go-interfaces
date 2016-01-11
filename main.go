package main

import (
	"fmt"
	"strconv"
)

//First Object, Square
type Square struct {
	side int
}

func (s *Square) area() string {
	return strconv.Itoa(s.side * s.side)
}

//Second Object Circle
type Circle struct {
	radius float32
}

func (s *Circle) area() string {
	return fmt.Sprintf("%f", 3.14*s.radius*s.radius)

}

//Interface for shape, which creates a string
type Shape interface {
	area() string
}

func printArea(s Shape) string {
	return fmt.Sprintf("area of square = %s \n", s.area())
}

func main() {
	fmt.Println("hello world")

	testSquare := &Square{side: 5}
	testCircle := &Circle{radius: 5}
	// fmt.Println(testSquare)

	fmt.Printf("area of square = %v \n", testSquare.area())

	fmt.Println(printArea(testSquare))
	fmt.Println(printArea(testCircle))

}
