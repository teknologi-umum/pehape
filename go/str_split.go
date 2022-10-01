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

	current := 0
	for i := 0; i < lenS; i += leng {
		if i < lenS {
			results = append(results, s[current:i])
			current += leng
		}
	}

	if mod > 0 {
		results = append(results, s[lenS-mod:])
	}
	return nil
}

func getCapacity(lenS, n int) (int, int) {
	mod := lenS % n
	if mod == 0 {
		return lenS / n, 0
	}

	return lenS/n + 1, mod
}
