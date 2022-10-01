package pehape

import "fmt"

func StrSplit(s string, length ...int) []string {
	lenS := len(s)
	if lenS == 0 {
		return []string{}
	}
	if len(length) > 1 {
		panic(fmt.Sprintf("expects at most 2 parameters, %d given", len(length)+1))
	}

	leng := 1
	if len(length) == 1 {
		leng = length[0]
	}

	if leng < 1 {
		panic("The length of each segment must be greater than zero")
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
