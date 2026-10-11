// Write a function to split a string of space-separated words and convert it into an array of words. Words in the input will be separated by exactly one space. The input will not have leading or trailing spaces. A "word" is any contiguous sequence of characters that does not contain any space.
// Examples (Input ==> Output):

// "string" ==> ["string"]
// "Robin Singh" ==> ["Robin", "Singh"]
// "I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

package kata

import (
	"strings"
)

func StringToArray(s string) []string {
	return strings.Split(s, " ") // Use the Split function from the strings package to split the input string by spaces and return the resulting slice of words

}
