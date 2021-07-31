package main

import "testing"

func BenchmarkWithoutPipelineModule(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		addFoo(addQuote(square(multiplyTwo(i))))
	}
}