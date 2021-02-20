package main

import "fmt"

type contactInfo struct {
	phone   string
	zipCode int
}

type person struct {
	firstname   string
	surname     string
	age         int
	contactInfo contactInfo // Embedded Stuct: contactInfo
}

func main() {
	//declaring, initialising, assigning Alex to person
	alex, paul, john := generatePersonSimple()

	fmt.Println(alex)
	fmt.Println(paul)
	fmt.Println(john)

	//generate person with "Zero Values": Struct Option 3
	generatePersonZeroValues()

	//call updateName with person
	john.updateName("Joe")
}

//Update firstname in person struct function that takes person receiver and a name<string> as arg
func (p person) updateName(newFirstname string) {
	p.firstname = newFirstname
}

// Simply generate Person Options
func generatePersonSimple() (person, person, person) {
	//Option 1: Go interprets this in the order of declared variables in struct
	alex := person{"Alex", "Byrne", 32, contactInfo{}}

	//Option 2: Create person by specifying corresponding vars
	paul := person{firstname: "Paul", surname: "Connor", age: 23}

	//Creating person with embedded struct
	john := person{
		firstname: "John",
		surname:   "Doe",
		age:       55,
		contactInfo: contactInfo{
			phone:   "04244342",
			zipCode: 95000,
		},
	}

	//return persons
	return alex, paul, john
}

//func generate Person Option Three
func generatePersonZeroValues() {

	//Create var Alex of type person: no assigned values yet
	//Thus, Go assigns a "Zero Value" to fields within person struct

	// --------------------------------------------
	// 	Type				||	Zero Value
	// --------------------------------------------
	// 	String				||		""
	// 	int					||		0
	// 	float				|| 		0
	// 	boolean				|| 		false
	// --------------------------------------------
	var alex person
	fmt.Println(alex)

	//%+v will print out each field name + value in Alex
	fmt.Printf("%+v", alex)

}
