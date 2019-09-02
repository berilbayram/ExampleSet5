package main

import "fmt"

type person struct {
	first string
	last string
	favFlavors []string
}

type vehicle struct {
	color string
	doors int
}

type truck struct {
	vehicle
	fourWheel  bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println("Example 1")
	fmt.Println()
	personOne := person{
		first: "Beril",
		last: "Bayram",
		favFlavors: []string {
			"chocolate",
			"peach",
			"banana",
		},
	}

	personTwo := person {
		first: "John",
		last: "Doe",
		favFlavors: []string {
			"caramel",
			"lemon",
			"hazelnut",
		},
	}
	fmt.Println(personOne.first)
	fmt.Println(personOne.last)

	for i,v := range personOne.favFlavors {
		fmt.Println(i,v)
	}

	fmt.Println(personTwo.first)
	fmt.Println(personTwo.last)

	for i,v := range personTwo.favFlavors {
		fmt.Println(i,v)
	}

	fmt.Println("Example 2")
	fmt.Println()

	m := map[string]person{
		personOne.last: personOne,
		personTwo.last: personTwo,
	}
	for _,v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("--------")
	}

	fmt.Println("Example 3")
	fmt.Println()

	t := truck{
		vehicle:   vehicle{
			color: "Black",
			doors: 4,
			},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			color: "Red",
			doors: 3,
		},
		luxury:  false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(s.color)
	fmt.Println(t.color)

	fmt.Println("Example 4")
	fmt.Println()


	personThree := struct {
		first     string
		last	  string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		last: "Bond",
		friends: map[string]int{
			"Money Penny":		 1,
			"John Doe":          2,
			"Jane Doe":          3,
		},
		favDrinks: []string{
			"Beer",
			"Water",
			"Martini",
		},
	}
	fmt.Println(personThree.first)
	fmt.Println(personThree.last)
	fmt.Println(personThree.friends)
	fmt.Println(personThree.favDrinks)

	for k, v := range personThree.friends {
		fmt.Println(k, v)
	}

	for i, val := range personThree.favDrinks {
		fmt.Println(i, val)
	}
}
