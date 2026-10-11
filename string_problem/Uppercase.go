// Write a function which converts the input string to uppercase.
//problem from codewars.com

package kata

func MakeUpperCase(str string) string {
    //var i int = 0;
    arrayLength := len(str); //store the size of input string
  
    inputString := []byte(str) //convert string to byte array for element wise operation
  
    for i := 0 ; i < arrayLength; i++{
        if (inputString[i] >= 'a' && inputString[i] <= 'z'){ //check if the character is lowercase
          inputString[i] = inputString[i] - 32;}
    }
  
    str = string(inputString) //convert byte array back to string
    return str
}