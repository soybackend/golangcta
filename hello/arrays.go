package main

import "fmt"

func main() {
	arr := [3]int {1, 2, 3}
	fmt.Println(arr)

	var sli []int
	sli = append(sli, 2)
	fmt.Println(sli)

	names := map[int] string {
		1 : "Jose Alvarez",
		2 : "Congreso de Tecnologias Abiertas",
	}
	fmt.Println(names)
}