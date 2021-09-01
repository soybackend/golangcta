package main

import "fmt"

type Cube struct {
	Width float64
	Height float64
	Depth float64
}

func main()  {
	c1 := Cube{}
	c1.Width = 2
	c1.Height = 3
	c1.Depth = 4

	var c2 Cube
	c2.Height = 2
	c2.Width = 3
	c2.Depth = 4

	c3 := Cube{
		Width: 1,
		Height: 2,
		Depth: 3,
	}

	fmt.Println(c3.Width)
}