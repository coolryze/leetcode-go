package validanagram

func isAnagram(s string, t string) bool {
	records := [26]int{}

	for _, ch := range s {
		records[ch-'a']++
	}
	for _, ch := range t {
		records[ch-'a']--
	}

	return records == [26]int{}
}
