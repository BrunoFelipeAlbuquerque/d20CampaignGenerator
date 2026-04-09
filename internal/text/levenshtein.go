package text

func Levenshtein(a, b string) int {
	la := len(a)
	lb := len(b)

	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,      // Delete
				dp[i][j-1]+1,      // Insert
				dp[i-1][j-1]+cost, // Replace
			)
		}
	}

	return dp[la][lb]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
