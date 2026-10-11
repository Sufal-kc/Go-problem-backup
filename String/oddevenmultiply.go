package kata
// This kata is about multiplying a given number by eight if it is an even number and by nine otherwise.

package kata

func SimpleMultiplication(n int) int {
  if ((n%2)==0){
    n = n*8
  }else{
    n = n* 9
  }
  return n // your code here
}