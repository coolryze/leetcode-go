package spiralmatrixii

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	startI, startJ := 0, 0
	offset := 1
	loopCount := n / 2
	count := 1
	for loopCount > 0 {
		for j := startJ; j < n-offset; j++ {
			res[startI][j] = count
			count++
		}
		for i := startI; i < n-offset; i++ {
			res[i][n-offset] = count
			count++
		}
		for j := n - offset; j > startJ; j-- {
			res[n-offset][j] = count
			count++
		}
		for i := n - offset; i > startI; i-- {
			res[i][startJ] = count
			count++
		}
		startI++
		startJ++
		offset++
		loopCount--
	}

	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}
	return res
}
