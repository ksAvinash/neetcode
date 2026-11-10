func majorityElement(nums []int) int {
	l := len(nums)
	counts := map[int]int{} // val: count
	for _, n := range nums {
		counts[n]++
		if counts[n] > l/2 {
			return n
		}
	}
	return 0
}
