// Given a string of digits, you should replace any digit below 5 with '0' and any digit 5 and above with '1'. Return the resulting string.

// Note: input will never be an empty string
package kata

func FakeBin(x string) string {
	// your code here
	//     eg: x:' 43563'
	result := []byte(x)
	n := len(result)
	for i := 0; i < n; i++ {
		if result[i] < '5' {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}
