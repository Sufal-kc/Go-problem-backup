package kata

import (
  "strings"
)

func MakeUpperCase(str string) string {
    return strings.ToUpper(str)
}
//imports from the strings package and uses the built-in ToUpper function to convert the input string to uppercase. This approach is more concise and efficient compared to manually iterating through each character in the string.