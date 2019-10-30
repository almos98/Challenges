/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	// switch {
	// case n == 0:
	// 	return 1
	// case n == 1:
	// 	return x
	// case n < 0:
	// 	return 1/myPow(x,-n)
	// default:
	// 	return x*myPow(x, n-1)
	// }

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	res := float64(1)
	for i := n; i > 0; i-- {
		res *= x
	}

	if neg {
		return 1/res
	}
	return res
}
// @lc code=end
