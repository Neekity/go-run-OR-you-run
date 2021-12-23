package prime

func PrimeA(n int) (res []int) {
	origin := make(chan int)
	go pushNumber2Channel(origin, n)
	for {
		next := make(chan int)
		prime, ok := <-origin
		if !ok {
			return
		}
		res = append(res, prime)
		go func(prime int, origin chan int, next chan int) {
			for i := range origin {
				if i%prime != 0 {
					next <- i
				}
			}
			close(next)
		}(prime, origin, next)
		origin = next
	}
}

func PrimeB(n int) (res []int) {
	origin := make(chan int)
	wait := make(chan struct{})
	go pushNumber2Channel(origin, n)

	go recursionProcessor(origin, wait, &res)
	<-wait
	return
}

func PrimeS(n int) (res []int) {
	filter := make([]int, 0)
	for prime := 2; prime < n; prime += 1 {
		if filterPrime(prime, filter) {
			filter = append(filter, prime)
			res = append(res, prime)
		}
	}
	return
}

func pushNumber2Channel(origin chan int, n int) {
	for num := 2; num < n; num++ {
		origin <- num
	}
	close(origin)
}

func recursionProcessor(origin chan int, wait chan struct{}, res *[]int) {
	prime, ok := <-origin
	if !ok {
		close(wait)
		return
	}
	*res = append(*res, prime)
	next := make(chan int)
	go recursionProcessor(next, wait, res)
	for num := range origin {
		if num%prime != 0 {
			next <- num
		}
	}
	close(next)
}

func filterPrime(x int, filter []int) bool {
	for _, f := range filter {
		if x < f*f {
			return true
		}
		if x%f == 0 {
			return false
		}
	}
	return true
}
