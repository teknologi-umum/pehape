package pehape

func StrSplit(s string, length ...int) []string {
	lenS := len(s)
	if lenS == 0 {
		return []string{}
	}
	if len(length) > 1 {
		panic("length's len must be 1")
	}

	leng := 1
	if len(length) == 1 {
		leng = length[0]
	}

	capacity, mod := getCapacity(lenS, leng)

	results := make([]string, 0, capacity)

	to := leng
	for i := 0; i < lenS; i += leng {
		if to <= lenS {
			results = append(results, s[i:to])
			to += leng
		}
	}

	if mod > 0 {
		results = append(results, s[lenS-mod:])
	}
	return results
}

func getCapacity(lenS, n int) (int, int) {
	mod := lenS % n
	if mod == 0 {
		return lenS / n, 0
	}

	return lenS/n + 1, mod
}
