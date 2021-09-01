package main

import (
	"fmt"
	"sync"
)

func main()  {
	const nt = 10
	wg := sync.WaitGroup{}
	wg.Add(nt)

	for i := 0; i < nt; i++ {
		numTask := i
		go func() {
			defer wg.Done()
			fmt.Println("Running tasks ", numTask)
		}()
	}

	wg.Wait()
	fmt.Println("All tasks were executed")
}
