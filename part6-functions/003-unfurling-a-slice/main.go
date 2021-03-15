package main

import "fmt"

// "展開"

func main(){
	xi := []int{2,3,4,5,6,7,8,9}
	x := sum(xi...)
	fmt.Println("total: ", x)
}
// slice
// passing zero or more
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x{
		fmt.Println("idx:", i)
		sum += v
	}
	return sum
}