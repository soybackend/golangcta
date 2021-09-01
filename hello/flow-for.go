package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	rand.Seed(time.Now().UnixNano())
	numero := rand.Intn(100)

	for numero % 2 == 0 {
		fmt.Println("El número ", numero, " es par")
		numero = rand.Intn(100)
	}
	fmt.Println("El número ", numero, " es impar")


	fmt.Println("Otro for")
	for {
		numero = rand.Intn(100)
		if numero % 2 != 0 {
			break
		}
		fmt.Println("El número del otro for ", numero, " es par")
	}
	fmt.Println("El número del otro for ", numero, " es impar")
}
