package pipeline

import (
	"testing"
)

func BenchmarkWithoutPipelineModule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addFoo(addQuote(square(multiplyByTwo(i))))
	}
}

func BenchmarkWithPipelineModule(b *testing.B) {
	outC := New(func(inc chan interface{}) {
		defer close(inc)
		for i := 0; i < b.N; i++ {
			inc <- i
		}
	}).Pipe(func(i interface{}) (interface{}, error) {
		return multiplyByTwo(i.(int)),nil
	}).Pipe(func(i interface{}) (interface{}, error) {
		return square(i.(int)),nil
	}).Pipe(func(i interface{}) (interface{}, error) {
		return addQuote(i.(int)),nil
	}).Pipe(func(i interface{}) (interface{}, error) {
		return addFoo(i.(string)),nil
	}).Merge()

	for _ = range outC {
		//fmt.Println(o)
	}
}