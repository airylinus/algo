package algorithm

// BinarySearch find index of niddle in haystack return -1 if not exist
func BinarySearch(haystack []int64, niddle int64) (idx int) {
	idx = -1
	left := 0
	right := len(haystack) - 1
	for (right - left) > 1 {
		middle := (right + left) / 2
		if haystack[middle] == niddle {
			idx = middle
			return
		}
		if haystack[middle] > niddle {
			right = middle
		} else {
			left = middle
		}
	}
	if right-left == 1 {
		if haystack[right] == niddle {
			idx = right
		}
		if haystack[left] == niddle {
			idx = left
		}
	}
	return
}
