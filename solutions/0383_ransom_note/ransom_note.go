package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	records := [26]int{}

	for _, ch := range ransomNote {
		records[ch-'a']++
	}
	for _, ch := range magazine {
		if records[ch-'a'] > 0 {
			records[ch-'a']--
		}
	}

	return records == [26]int{}
}
