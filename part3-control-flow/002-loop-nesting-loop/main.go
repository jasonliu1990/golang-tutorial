package main

import "fmt"

func main()  {
	for i := 1; i <= 5; i++{
		for j := 1; j <= 3; j++{
			k := i * j
			fmt.Printf("%d x %d = %d \n", i, j, k)
		}
	}
}
