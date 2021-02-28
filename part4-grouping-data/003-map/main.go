package main

import "fmt"

func main() {
	// 1. introduction
	fmt.Println("1. introduction")
	// map[key]value{}
	m := map[string]int{
		"James":      32,
		"Moneypenny": 26,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Moneypenny"])
	fmt.Println(m["???"]) // 0

	v, ok := m["???"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["???"]; ok{
		fmt.Println("this is the if print", v)
	}
	fmt.Println("=========================")
	// 2. add element & range
	fmt.Println("2. add element & range")
	// add element
	m["Jaz"] = 31
	fmt.Println(m)
	// range
	for k, v := range m{
		fmt.Println(k, v)
	}
	// compare w/
	xi := []int{4, 5, 6, 7, 8}
	for i, v := range xi{
		fmt.Println(i, v)
	}
	fmt.Println("=========================")
	// 3. delete
	fmt.Println("3. delete")
	fmt.Println(m)
	delete(m, "Jaz")
	fmt.Println(m)
	delete(m, "???") // 不存在的也不會報錯
	// 如何確認刪除的東西存在?
	if v, ok := m["James"]; ok{
		fmt.Println("value:", v)
		delete(m, "James")
	}
	fmt.Println(m)
	fmt.Println("=========================")

}
