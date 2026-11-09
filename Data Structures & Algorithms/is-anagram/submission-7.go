// import "strings"

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(s) != len(t) {
		return false
	}
	sCounts, tCounts := [26]rune{}, [26]rune{}
	for _, c := range s {
		sCounts[c-'a']++
	}
	for _, c := range t {
		tCounts[c-'a']++
	}
	// fmt.Println(sCounts, tCounts)
	return sCounts == tCounts
}
