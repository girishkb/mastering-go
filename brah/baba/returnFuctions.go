package main

import "fmt"

func main()  {
	i := funcReturn()
	j := funcReturn()
	fmt.Println(i(),j())
	fmt.Println(i(),j())
	fmt.Println(i(),j())

}

func funcReturn() func() int {
	i :=0
	return func() int {
		i++
		return i*i
	}
}