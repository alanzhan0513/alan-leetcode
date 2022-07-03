func findKthLargest(nums []int, k int) int {
    return QuickSelect(nums, 0, len(nums) - 1, k)
}

func QuickSelect(arr []int, front int, end int, k int) int {
	if len(arr) == 0 || len(arr) < k {
		return -1
	}

	if front == end {
		return arr[front]
	}

	pivotIdx := partition(arr, front, end)
	for pivotIdx != k - 1 {
		if pivotIdx > k - 1 {
			return QuickSelect(arr, front, pivotIdx - 1, k);
		} else {
			return QuickSelect(arr, pivotIdx + 1, end, k)
		}
	}

	return arr[pivotIdx]
}

func partition(arr []int, front int, end int) int {
	pivot := arr[end]
	i := front - 1
	for j := front; j < end; j++ {
		if arr[j] > pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	i++
    arr[i], arr[end] = arr[end], arr[i]
	return i
}
