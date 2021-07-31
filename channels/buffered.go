package main

import (
	"fmt"
)

//func main()  {
//	/*ch := make(chan string,2)
//	ch <- "girish"
//	ch <- "giddy"
//	//close(ch)
//	for c := range  ch {
//		fmt.Println(c)
//	}*/
//
//	ch := make(chan int,2)
//	go write(ch)
//	time.Sleep(20*time.Second)
//	for c := range  ch {
//		fmt.Println("read value ", c , " from chan")
//		time.Sleep(5*time.Second)
//	}
//}

func write(ch chan int)  {
	for i := 0; i < 5 ; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
