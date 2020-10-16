package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	args := os.Args
	if len(args) != 2  {
		fmt.Println("The program need 1 arguments")
		return
	}
	y,err :=	strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	square := func(n int) int{
		return n *n
	}
	fmt.Println("Square of y ",y, "is ",square(y))
	double := func(s int) int{
		return s+s
	}

	fmt.Println("Double of y ",y, "is ",double(y))
}
