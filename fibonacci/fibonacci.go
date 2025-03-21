package fibonacci

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-2) + fib(n-1)
}

func fibTCO(n, a, b int) (nn, aa, bb int) {
	switch n {
	case 0:
		return a, a, b
	case 1:
		return b, a, b
	default:
		return fibTCO(n-1, b, a+b)
	}
}

func fibIter(n int) int {
	a, b := 0, 1
	for n > 1 {
		a, b = b, a+b
		n--
	}
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return b
	}
}
