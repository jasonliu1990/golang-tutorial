package main

import "fmt"

func main()  {
	fmt.Println(4*3*2*1)
	n := factorial(4)
	fmt.Println(n)

	n2 := f2(4)
	fmt.Println(n2)
}
// 遞迴
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	// 4 * f(4-1) * f(3-1) * f(2-1) * 1
 }

 func f2(n int) int{
 	ans := 1
	for ; n > 0; n--{
		ans *= n
	}
	return ans
 }