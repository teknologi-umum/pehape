package main

//make implode like php implode
func Implode(separator string, array []string) string {
	var result string

	for i := 0; i < len(array); i++ {
		if i == 0 {
			result = array[i]
		} else {
			result = result + separator + array[i]
		}
	}
	return result
}

//make implode like php with 2 params, the first param should be string or null, and the second param should be array
func Implode2(separator string, array []string) string {
	var result string

	if separator == "" {
		for i := 0; i < len(array); i++ {
			if i == 0 {
				result = array[i]
			} else {
				result = result + array[i]
			}
		}
	} else {
		for i := 0; i < len(array); i++ {
			if i == 0 {
				result = array[i]
			} else {
				result = result + separator + array[i]
			}
		}
	}
	return result
}