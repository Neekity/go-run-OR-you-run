package whynotcolumnfirst

func CopyRowFirst4096() (dst [4096][4096]bool) {
	src := generateData4096()
	for i := 0; i < 4096; i++ {
		for j := 0; j < 4096; j++ {
			dst[i][j] = src[i][j]
		}
	}
	return
}

func CopyColumnFirst4096() (dst [4096][4096]bool) {
	src := generateData4096()
	for j := 0; j < 4096; j++ {
		for i := 0; i < 4096; i++ {
			dst[i][j] = src[i][j]
		}
	}
	return
}

func generateData4096() (src [4096][4096]bool) {
	for i := 0; i < 4096; i++ {
		for j := 0; j < 4096; j++ {
			src[i][j] = true
		}
	}
	return
}
