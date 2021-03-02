package main

import "fmt"

// ex1
//Create your own type “person” which will have an underlying type of “struct”
//so that it can store the following data:
//	first name
//	last name
//	favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values,
//ranging over the elements in the slice which stores the favorite flavors.
type person struct {
	first  string
	last   string
	flavor []string
}

func ex1() {
	p1 := person{
		first:  "jaz",
		last:   "liu",
		flavor: []string{"coke", "whisky"},
	}

	p2 := person{
		first:  "haha",
		last:   "liu",
		flavor: []string{"tea", "water"},
	}

	fmt.Println(p1.first, p1.last)
	for i, f := range p1.flavor {
		fmt.Println(i, f)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.flavor {
		fmt.Println(i, v)
	}
}

// ex2
//Take the code from the previous exercise,
//then store the values of type person in a map with the key of last name.
//Access each value in the map. Print out the values, ranging over the slice.

func ex2() {
	p1 := person{
		first:  "jaz",
		last:   "liu",
		flavor: []string{"coke", "whisky"},
	}

	p2 := person{
		first:  "haha",
		last:   "liu",
		flavor: []string{"tea", "water"},
	}
	personMap := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for _, v := range personMap {
		fmt.Println(v.first, v.last, ":")
		for _, v2 := range v.flavor {
			fmt.Println(v2)
		}
	}
}

// ex3
// a. Create a new type: vehicle.
//		1. The underlying type is a struct.
//		2. The fields:
//			-doors
//			-color
// b. Create two new types: truck & sedan.
//		1. The underlying type of each of these new types is a struct.
//		2. Embed the “vehicle” type in both truck & sedan.
//		3. Give truck the field “fourWheel” which will be set to bool.
//		4. Give sedan the field “luxury” which will be set to bool.
type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

// c. Using the vehicle, truck, and sedan structs:
//		1. using a composite literal, create a value of type truck and assign values
//	       to the fields;
//		2. using a composite literal, create a value of type sedan and assign values
//	   	   to the fields.
// d. Print out each of these values.
// e. Print out a single field from each of these values.

func ex3() {
	v1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(v1, v2)
	fmt.Println(v1.color, v2.color)
	fmt.Println(v1.doors, v2.doors)

}

// ex4
// Create and use an anonymous struct.
func ex4() {
	v1 := struct {
		doors int
		color string
	}{
		doors: 4,
		color: "green",
	}

	fmt.Println(v1.doors, v1.color)
}

func main() {
	fmt.Println("======ex1======")
	ex1()
	fmt.Println("======ex2======")
	ex2()
	fmt.Println("======ex3======")
	ex3()
	fmt.Println("======ex4======")
	ex4()
}
