package main

import (
	"fmt"
	"sync"
	"testing"
)

const (
	numWorkers = 1000
	numTasks   = 10000
)

func workWithPool(pool *sync.Pool) {
	for i := 0; i < numTasks; i++ {
		obj := pool.Get()
		// Simulate some work with the object
		_ = obj
		pool.Put(obj)
	}
}

func workWithWaitGroup(wg *sync.WaitGroup) {
	for i := 0; i < numTasks; i++ {
		// Simulate some work
	}
	wg.Done()
}

func workWithChannel(ch chan struct{}) {
	for i := 0; i < numTasks; i++ {
		// Simulate some work
	}
	ch <- struct{}{}
}

func BenchmarkPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() interface{} {
			return new(interface{})
		},
	}
	b.ReportAllocs()
	b.ResetTimer()

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
	b.ReportAllocs()
	b.ResetTimer()

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
	b.ReportAllocs()
	b.ResetTimer()

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

func main() {

	// sync.Pool oluşturuyoruz ve bağlantı nesnelerini tutacak.
	pool := &sync.Pool{
		New: func() interface{} {
			println("New Creation")
			return &Connection{}
		},
	}

	// İstemci istekleri için bir döngü simüle edelim.
	for i := 0; i < 10; i++ {
		// Havuzdan bir bağlantı alıyoruz veya havuzda yoksa yeni bir bağlantı oluşturulacak.
		conn := pool.Get().(*Connection)
		conn.ID += i

		// İstemci taleplerine göre işlem yapıyoruz.
		// Burada sadece bağlantı numarasını ekrana yazdırıyoruz.
		fmt.Println("Handling request with connection ID:", conn.ID)

		// İşlem bittikten sonra bağlantıyı havuza geri bırakıyoruz.
		//pool.Put(conn)
	}

	/*	poolBenchmark := testing.Benchmark(BenchmarkPool)
		fmt.Printf("Pool: %s\n", poolBenchmark)

		wgBenchmark := testing.Benchmark(BenchmarkWaitGroup)
		fmt.Printf("WaitGroup: %s\n", wgBenchmark)

		channelBenchmark := testing.Benchmark(BenchmarkChannel)
		fmt.Printf("Channel: %s\n", channelBenchmark)*/
}

type Connection struct {
	ID int
}
