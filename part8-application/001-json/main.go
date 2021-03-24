package main

import (
	"encoding/json"
	"fmt"
)

type person struct{
	First string
	Last  string
	Age   int
}

func main()  {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   26,
	}

	people := []person{p1, p2}
	//fmt.Println(people)

	//marshal
	bs, err := json.Marshal(people)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	//unmarshal
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":26}]
`
	bs2 := []byte(s)
	var people2 []person

	err = json.Unmarshal(bs2, &people2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(people2)
	for i, v := range people2{
		fmt.Println(i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}