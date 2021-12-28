package damnifelse

var src [500]int

func init() {
	for i := 0; i < 500; i++ {
		src[i] = i % 300
	}
}

func CountGtNumberElse() (res int) {
	res = 0
	for count := 1; count < 100000; count++ {
		for i := 0; i < 500; i++ {
			if src[i] < 260 {
				res -= 1
			} else {
				res += 1
			}
		}
	}

	return res
}

func CountGtNumberIf() (res int) {
	res = 0
	for count := 1; count < 100000; count++ {
		for i := 0; i < 500; i++ {
			switch {
			case src[i] >= 260:
				res -= 1
			case src[i] < 260:
				res += 1
			}
		}
	}
	return res
}
