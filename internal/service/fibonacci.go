package service

func (fs *fibonacciService) GetFibonacci(n int64) int64 {
	var i, first, second int64

	if n < 2 {
		return n
	}

	first, second = 0, 1
	for i = 2; i < n; i++ {
		tmp := first + second
		first = second
		second = tmp
	}
	return first + second
}
