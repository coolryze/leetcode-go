package reversestringii

func reverseStr(s string, k int) string {
	bs := []byte(s)
	l := len(s)

	for i := 0; i < l; i += 2 * k {
		if i+k <= l {
			reverseString(bs[i : i+k])
		} else {
			reverseString(bs[i:l])
		}
	}

	return string(bs)
}

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
