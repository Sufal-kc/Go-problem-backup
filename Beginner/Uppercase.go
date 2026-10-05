// Write a function which converts the input string to uppercase.
//problem from codewars.com
package kata

func MakeUpperCase(str string) string {
    //var i int = 0;
    arrayLength := len(str);
  
    inputString := []byte(str)
  
    for i := 0 ; i < arrayLength; i++{
        if (inputString[i] >= 'a' && inputString[i] <= 'z'){
          inputString[i] = inputString[i] - 32;}
    }
  
    str = string(inputString)
    return str
}