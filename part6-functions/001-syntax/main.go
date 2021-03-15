package main

import "fmt"

func main()  {
	foo()
	bar("james")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
}
// func (r receiver) identifier(parameters) (return(s)) {...}
func foo(){
	fmt.Println("foo")
}

// everything in go is pass by value
func bar(s string){
	fmt.Println("hello", s)
}

// return
func woo(s string) string{
	return fmt.Sprint("hello ", s)
}

func mouse(fn string, ln string)(string, bool){
	a := fmt.Sprint(fn, " " ,ln, " says hello")
	b := false
	return a, b
}