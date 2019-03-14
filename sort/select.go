// Package sort contains sort functions
// select sort
// time complexity: O (n^2)
package sort

// SelectSort return sorted haystack
func SelectSort(haystack []int64) (sorted []int64) {
	total := len(haystack)
	var n int64
	for len(sorted) < total {
		n, haystack = findBiggest(haystack)
		sorted = append(sorted, n)
	}
	return
}

func findBiggest(haystack []int64) (val int64, newHaystack []int64) {
	i := 0
	maxElementIdx := 0
	for i < len(haystack) {
		if haystack[i] > val {
			val = haystack[i]
			maxElementIdx = i
		}
		i++
	}
	newHaystack = append(haystack[:maxElementIdx], haystack[maxElementIdx+1:]...)
	return
}
