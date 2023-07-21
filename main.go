package main

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

	for data := range resultCh {
		result[data[0]-1] = data[1]
		println(data)
	}

	//close(resultCh)
	return result
}

func CalculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return CalculateFibonacci(n-1) + CalculateFibonacci(n-2)
}
