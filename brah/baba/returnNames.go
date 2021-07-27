package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("The program needs at least 2 arguments")
		return
	}
	x,_ :=	strconv.Atoi(args[1])
	y,_ :=	strconv.Atoi(args[2])
	fmt.Println(minMax(x,y))
	fmt.Println(namedMinMax(x,y))
}

func namedMinMax(x, y int)  (min,max int){
	if x > y {
		min = y
		max = x
	} else {
		max = y
		min = x
	}
	return
}

func minMax(x, y int)  (min,max int){
	if x > y {
		min = y
		max = x
	} else {
		max = y
		min = x
	}
	return min,max
}
