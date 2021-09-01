package main

import (
	"fmt"
	"hello/cube"
)

func Hola(nombre string, apellido string) {
	fmt.Printf("Hola, %s %s \n", nombre, apellido)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	Hola("Jose", "Alvarez")
	fmt.Println(max(5, 4))

	fmt.Println(cube.Volume(3))
}
