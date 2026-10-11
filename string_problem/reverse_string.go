// Complete the solution so that it reverses the string passed into it.

// 'world'  =>  'dlrow'
// 'word'   =>  'drow'

package kata

func ReverseString(str string) string{
	strArray := []rune(str) // Convert the string to a slice of runes to handle Unicode characters
	for i, j := 0, len(strArray)-1; i < j; i, j = i+1, j-1 {
		strArray[i], strArray[j] = strArray[j], strArray[i] // Swap the characters
	}
	return string(strArray) // Convert the slice of runes back to a string and return it
}

//or

package kata

func ReverseString(str string) string{
	  reversed := ""
	  for i := len(str) - 1; i >= 0; i-- {
		    reversed += string(str[i]) // Append each character in reverse order
	  }
	  return reversed
	  
}