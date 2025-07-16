package main

import (
	"fmt"
	"strings"
)

type Person struct {
	personName string
	personAge  int
}

// storing hardcoded values in the slice
var personSlice = []Person{
	{personName: "Advait", personAge: 28},
	{personName: "Atharv", personAge: 18},
	{personName: "Aaryan", personAge: 11},
	{personName: "Danish", personAge: 67},
	{personName: "pratik", personAge: 18},
}
//checking if the person is eligible for voting
func (person Person) checkVoteEligibility() bool {
	if person.personAge >= 18 {
		return true
	}
	return false
}
// fetching all person details with option to update the age
func personsIntroduction() {
	var name string
	fmt.Print("Enter the person name to see the details\n")
	fmt.Scan(&name)
	found := false
	for _, value := range personSlice {
		if strings.EqualFold(value.personName, name) {
			fmt.Printf("Hello!!  %s  you are  %d years old\n", value.personName, value.personAge)
			if value.checkVoteEligibility() == true {
				fmt.Println("you are eligible to vote")

			} else {
				fmt.Println("you are not eligible to vote")
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Person not found!! please enter valid name")
		personsIntroduction()
	}
	fmt.Println("Enter 'yes' to update age 'no' for main menu")
	var choice string
	fmt.Scan(&choice)
	if choice == "yes" {
		updateAge(name)
	} else if choice == "no" {
		return
	}
}

// updating the age of person and sending message if updated age is eligible for voting
func updateAge(name string) {
	var newAge int
	fmt.Print("Enter new age")
	fmt.Scan(&newAge)
	found := false
	for iterator := range personSlice {
		if strings.EqualFold(personSlice[iterator].personName, name) {
			priviousAge := personSlice[iterator].personAge
			personSlice[iterator].personAge = newAge
			fmt.Println("Your age has been updated successfully!")
			if priviousAge < 18 && newAge >= 18 {
				fmt.Println("You are now eligible for voting")
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Println("person not found")
	}

}
// func main() {

// 	personsIntroduction()
// }
