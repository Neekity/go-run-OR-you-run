package whynotcolumnfirst

func CopyRowFirst2048() (dst [2048][2048]bool) {
	src := generateData2048()
	for i := 0; i < 2048; i++ {
		for j := 0; j < 2048; j++ {
			dst[i][j] = src[i][j]
		}
	}
	return
}

func CopyColumnFirst2048() (dst [2048][2048]bool) {
	src := generateData2048()
	for j := 0; j < 2048; j++ {
		for i := 0; i < 2048; i++ {
			dst[i][j] = src[i][j]
		}
	}
	return
}

func generateData2048() (src [2048][2048]bool) {
	for i := 0; i < 2048; i++ {
		for j := 0; j < 2048; j++ {
			src[i][j] = true
		}
	}
	return
}
