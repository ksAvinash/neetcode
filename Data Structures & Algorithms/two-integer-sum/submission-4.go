func twoSum(nums []int, target int) []int {
    rem := map[int]int{}
	for i, n := range nums {
		r := target - n
		v, ex := rem[r]; if ex {
			return []int{v, i}
		}
		rem[n] = i
	}
	return []int{}
}
