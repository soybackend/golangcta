package main

import "fmt"

func cincoMensajes (msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("(%d de 5) %s\n", i, msg)
	}
}

func main()  {
	fmt.Println("Lanzando...")
	go cincoMensajes("Puede que no siempre termine...")
	cincoMensajes("Siempre termina.")
	fmt.Println("Finalizado.")
}
