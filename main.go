package main

import "time"

func main() {
	a := CallFibonacciFuncWithChannel(10)
	println(a)
}

func CallFibonacciFuncWithChannel(count int) []int {
	resultCh := make(chan []int, count)
	result := make([]int, count)
	for i := 1; i <= count; i++ {
		go func(index int) {
			resultCh <- []int{index, CalculateFibonacci(index)}
		}(i)
	}

	go func() {
		for data := range resultCh {
			result[data[0]-1] = data[1]
			println(data)
		}
	}()

	time.Sleep(time.Second * 5)

	return result
}

func CalculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return CalculateFibonacci(n-1) + CalculateFibonacci(n-2)
}
