package pipeline

import (
	"fmt"
	"time"
)

func multiplyByTwo(i int)  int {
	time.Sleep(1*time.Second)
	return i*2
}

func square(i int) int {
	time.Sleep(1*time.Second)
	return i*i
}

func addQuote(v int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("'%d'", v)
}
func addFoo(v string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}