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