// Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).
// Examples

// n = 0  ==> [1]        # [2^0]
// n = 1  ==> [1, 2]     # [2^0, 2^1]
// n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]

package kata

// import "math"

func PowersOfTwo(n int) []uint64 {
	// your code goes here
	power := []uint64{}
	var x uint64 = 1

	for i := 0; i <= n; i++ {
		power = append(power, x)
		x = x * 2

	}
	return power
}

//or
// package kata

func PowersOfTwobit(n int) (arr []uint64) {
	for i := 0; i <= n; i++ {
		arr = append(arr, 1<<i) // 001 =1 , 010 =2, 100=4, 1000=8, 10000=16, 100000=32, 1000000=64, 10000000=128, 100000000=256, 1000000000=512
	}
	return
}
