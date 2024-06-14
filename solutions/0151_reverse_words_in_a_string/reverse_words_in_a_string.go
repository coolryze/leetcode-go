package reversewordsinastring

func reverseWords(s string) string {
	strBytes := []byte(s)

	slow := 0
	for fast := 0; fast < len(strBytes); fast++ {
		if strBytes[fast] != ' ' {
			if slow != 0 {
				strBytes[slow] = ' '
				slow++
			}

			for fast < len(strBytes) && strBytes[fast] != ' ' {
				strBytes[slow] = strBytes[fast]
				slow++
				fast++
			}
		}
	}
	strBytes = strBytes[:slow]
	reverseString(strBytes)

	last := 0
	for i := 0; i <= len(strBytes); i++ {
		if i == len(strBytes) || strBytes[i] == ' ' {
			reverseString(strBytes[last:i])
			last = i + 1
		}
	}

	return string(strBytes)
}

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
