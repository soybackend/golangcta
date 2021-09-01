package main

import "fmt"

func main()  {
	i := 10
	p := &i
	a := p
	*p = 21
	*a = 50

	fmt.Println("El valor de a es", a)
	fmt.Println("El valor de i es", i)
	fmt.Println("El valor de p es", p)
	fmt.Println("El valor de la variable apuntada por p (o sea a) es", *p)
}
