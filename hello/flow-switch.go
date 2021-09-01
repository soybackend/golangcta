package main

import (
	"fmt"
	"runtime"
)

func main()  {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS x")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println(os)
	}
}