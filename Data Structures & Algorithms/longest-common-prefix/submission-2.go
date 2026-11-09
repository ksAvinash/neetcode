func longestCommonPrefix(strs []string) string {
    res := ""
	for i, c := range strs[0] { // b a t
		for _, s := range strs {
			if i >= len(s) || rune(s[i]) != c {
				return res
			}
		}
		res += string(c)
	}
	return res
}
