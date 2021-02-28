package main

import "fmt"

func main() {
	// 1. composite literal
	// x := type{values}
	fmt.Println("1. composite literal")
	x := []int{4, 5, 6, 7, 8, 42} //group together values of the same type
	fmt.Println(x)
	fmt.Println("=========================")

	// 2. for range
	fmt.Println("2. for range")
	//x := []int{4, 5, 6, 7, 8, 42}
	for i, v := range x {
		fmt.Println(i, v)
	}
	// not use range
	for i := 0; i < len(x); i++{
		fmt.Println(i, x[i])
	}
	fmt.Println("=========================")

	// 3. slicing a slice
	fmt.Println("3. slicing a slice")
	//x := []int{4, 5, 6, 7, 8, 42}
	//
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println("=========================")

	// 4. append to a slice
	fmt.Println("4. append to a slice")
	//x := []int{4, 5, 6, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99)
	fmt.Println(x)
	y := []int{123, 456, 789}
	//... mean take something that has a whole bunch of values in a data
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("=========================")

	// 5. del from a slice
	fmt.Println("5. delete from a slice")
	//x := []int{4, 5, 6, 7, 8, 42}
	// 刪除 x[3]
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
	fmt.Println("=========================")

	// 6. make
	fmt.Println("6. make")
	//x := []int{4, 5, 6, 7, 8, 42}
	x6 := make([]int, 10, 11)
	fmt.Println(x6)
	fmt.Println(len(x6))
	fmt.Println(cap(x6))
	x6[0] = 42
	x6[9] = 999
	fmt.Println(x6)
	fmt.Println(len(x6))
	fmt.Println(cap(x6))
	x6 = append(x6, 1234)
	fmt.Println(x6)
	fmt.Println(len(x6))
	fmt.Println(cap(x6))
	// 超過時會自動配置長度
	// double size of that
	x6 = append(x6, 1235)
	fmt.Println(x6)
	fmt.Println(len(x6))
	fmt.Println(cap(x6))
	fmt.Println("=========================")

	// 7. multi-dimensional slice
	fmt.Println("7. multi-dimensional slice")
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "penny", "strawberry", "hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	fmt.Println("=========================")
}
