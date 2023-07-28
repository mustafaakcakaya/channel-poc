package main

import "fmt"

func main() {
	a := CallFibonacciFuncWithChannel(10)
	fmt.Printf("%v\n", a)
}

func CallFibonacciFuncWithChannel(count int) []int {
	resultCh := make(chan []int, count)
	result := make([]int, count)
	for i := 1; i <= count; i++ {
		go func(index int) {
			resultCh <- []int{index, CalculateFibonacci(index)}
		}(i)
	}

	for i := 0; i < count; i++ {
		data := <-resultCh
		result[data[0]-1] = data[1]
	}
	close(resultCh)
	return result
}

func CalculateFibonacci(n int) int {
	if n <= 1 || n > 10 {
		return n
	}
	return CalculateFibonacci(n-1) + CalculateFibonacci(n-2)
}
