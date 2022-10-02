package pehape

// Levenshtein calculates the distance between two strings.
// This algorithm allow insertions, deletions and substitutions to change one string to the second
// Compatible with non-ASCII characters
func Levenshtein(str1 string, str2 string, cost ...int) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	insertionCost := 1
	replacementCost := 1
	deletionCost := 1
	if len(cost) >= 1 {
		insertionCost = cost[0]
	}
	if len(cost) >= 2 {
		replacementCost = cost[1]
	}
	if len(cost) >= 3 {
		deletionCost = cost[2]
	}

	// Get and store length of these strings
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	if runeStr1len == 0 {
		return runeStr2len
	} else if runeStr2len == 0 {
		return runeStr1len
	} else if equal(runeStr1, runeStr2) {
		return 0
	}

	column := make([]int, runeStr1len+1)

	for y := 1; y <= runeStr1len; y++ {
		column[y] = y
	}
	for x := 1; x <= runeStr2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= runeStr1len; y++ {
			oldkey := column[y]
			var i int
			if runeStr1[y-1] != runeStr2[x-1] {
				i = replacementCost
			}
			column[y] = min(
				min(column[y]+insertionCost, // insert
					column[y-1]+deletionCost), // delete
				lastkey+i) // substitution
			lastkey = oldkey
		}
	}

	return column[runeStr1len]
}

// Return the smallest integer among the two in parameters
func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

// Compare two rune arrays and return if they are equals or not
func equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
