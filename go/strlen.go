package pehape

// Then Strlen() function calculates the length of the given string like php does.
// Parameter
//   - str => string to calculate the length
//
// Returns
//   - int => the length of the string
func Strlen(str string) int {
	return len(str)
}
