package main

import "fmt"

// 可變參數

func main(){
	x := sum(2,3,4,5,6,7,8,9)
	fmt.Println("total: ", x)
}
//slice
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