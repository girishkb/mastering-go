package main

import "fmt"

func main() {
	fmt.Println("function1 : ", funFunction(function1, 10))
	fmt.Println("function2 : ", funFunction(function2, 10))
	fmt.Println("Inline : ",
		funFunction(func(i int) int {
			return i * i * i
		}, 10))

}

func funFunction(f func(int) int, v int) int {
	return f(v)
}

func function1(x int) int {
	return x + x
}

func function2(x int) int {
	return x * x
}
