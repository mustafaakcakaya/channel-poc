package main

import (
	"sync"
	"testing"
)

func BenchmarkPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() interface{} {
			return new(interface{})
		},
	}
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(numWorkers)
		for i := 0; i < numWorkers; i++ {
			go func() {
				workWithPool(pool)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkWaitGroup(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(numWorkers)
		for i := 0; i < numWorkers; i++ {
			go func() {
				workWithWaitGroup(&wg)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkChannel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ch := make(chan struct{}, numWorkers)
		for i := 0; i < numWorkers; i++ {
			go func() {
				workWithChannel(ch)
			}()
		}
		// Wait for all workers to finish
		for i := 0; i < numWorkers; i++ {
			<-ch
		}
	}
}
