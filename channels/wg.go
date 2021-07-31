package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup)  {
	fmt.Println("Started Goroutine ", i)
	time.Sleep(5*time.Second)
	fmt.Println("Goroutine ended\n ", i)
	wg.Done()
}

func main()  {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i,&wg)
	}
	wg.Wait()
}
