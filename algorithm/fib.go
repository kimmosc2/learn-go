package algorithm

var fibCache = []int{}

// fib1 not memory
func fib1(num int) int {
	if num > 0 {
		if num < 2 {
			return 1
		}
		return fib1(num-1) + fib1(num-2)
	}
	return 0
}

// fib2 has cache from memory
func fib2(num int) int {
	if num > 0 {
		if fibCache[num] != 0 {
			return fibCache[num]
		}
		if num <= 1 {
			fibCache[num] = 1
			return 1
		}
		fibCache[num] = fib2(num-1) + fib2(num-2)
		return fibCache[num]
	}
	return 0
}
