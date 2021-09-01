package main

import (
	"fmt"
	"sync"
)

func main()  {
	const numTasks = 3
	wg := sync.WaitGroup{}
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		numTask := i
		go func() {
			defer wg.Done()
			fmt.Println("Running tasks ", numTask)
		}()
	}

	wg.Wait()
	fmt.Println("All tasks were executed")
}
