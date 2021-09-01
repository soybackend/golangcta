package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	valor := rand.Intn(100)
	fmt.Println(valor)

	if valor % 2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}
}
