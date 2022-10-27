func pageCount(n int32, p int32) int32 {
	// Write your code here
	if n%2 == 0 {
		n++
	}

	var turnFromBeg = p / 2.0
	var turnFromEnd = (n - p) / 2.0

	if turnFromBeg < turnFromEnd {
		return turnFromBeg
	} else {
		return turnFromEnd
	}
}
