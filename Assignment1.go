package main

import (
	"fmt"
)

type Person struct {
	personName string
	personAge  int
}

// storing structure in slice
var personSlice = []Person{}

// storing hardcoded values
var personData = []Person{
	{personName: "Advait", personAge: 28},
	{personName: "Atharv", personAge: 18},
	{personName: "Aryan", personAge: 11},
}

func (person Person) checkVoteEligibility() bool {
	return person.personAge >= 18
}

func personsIntroduction() {
	var name string
	fmt.Println("Enter the name to see the details")
	fmt.Scan(&name)
	for _, value := range personSlice {
		if value.personName == name {
			fmt.Printf("Hello  %s you are  %d years old\n", value.personName, value.personAge)
			if value.checkVoteEligibility() {
				fmt.Println("you are eligible to vote")
			} else {
				fmt.Println("you are not eligible to vote")
			}
		}
	}
	fmt.Println("please Enter 'yes' to update age")
	var choice string
	fmt.Scan(&choice)
	if choice == "yes" {
		updateAge(name)
	} else {
		return
	}
}

// function to update the age of person
func updateAge(name string) {
	var newAge int
	fmt.Print("Enter new age")
	fmt.Scan(&newAge)
	for iterator := range personSlice {
		if personSlice[iterator].personName == name {
			personSlice[iterator].personAge = newAge
          fmt.Println("Your age has been updated successfully!")
		}
	}

}
func main() {
	 personSlice = append(personSlice, personData...)
	personsIntroduction()
}
