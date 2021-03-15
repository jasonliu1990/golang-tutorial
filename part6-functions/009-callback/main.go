package main

import "fmt"
//回呼函式（callback function）是指能藉由參數（argument）通往另一個函式的函式。
//它會在外部函式內調用、以完成某些事情。
func main()  {
	ii := []int{1,2,3,4,5,6,7,8,9,10}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println(s2)

	s3 := odd(sum, ii...)
	fmt.Println(s3)
}

func sum(xi ...int)int{
	total := 0
	for _, v := range xi{
		total += v
	}
	return total
}
// callback
func even(f func(xi ...int) int, vi ...int) int{
	var yi []int
	for _, v := range vi{
		if v % 2 == 0{
			yi = append(yi, v)
		}
	}
	ans := f(yi...)
	return ans
}

func odd(f func(xi ...int) int, vi ...int) int{
	var yi []int
	for _, v := range vi{
		if v % 2 != 0{
			yi = append(yi, v)
		}
	}
	ans := f(yi...)
	return ans
}