func removeElement(nums []int, val int) int {
    v := 0
	for _, n := range nums {
		if n != val {
			nums[v] = n
			v++
		}
	}
	return v
}
