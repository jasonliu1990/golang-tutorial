package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	fmt.Println("======ex1======")
	ex1()
	fmt.Println("======ex2======")
	ex2()
}
// ex1
// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

func ex1(){

	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		fmt.Println("hello func 1")
		wg.Done()
	}()

	go func(){
		fmt.Println("hello func 2")
		wg.Done()
	}()
	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("gs", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("exit")
	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("gs", runtime.NumGoroutine())
}
// ex2
// This exercise will reinforce our understanding of method sets:
//		1. create a type person struct
//		2. attach a method speak to type person using a pointer receiver
//			*person
//		3. create a type human interface
//			to implicitly implement the interface,
//			a human must have the speak method
//		4. create func “saySomething”
//			have it take in a human as a parameter
//			have it call the speak method
//		5. show the following in your code
//			you CAN pass a value of type *person into saySomething
//			you CANNOT pass a value of type person into saySomething

