package main

import "fmt"

var a int
// we can create our own type in Go
type hotdog int
var b hotdog

func main(){
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b) // main.hotdog

	//a = b //cannot use b (type hotdog) as type int in assignment
	// how to make it work?
	// A: conversion, not casting
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
