package main

import (
	"errors"
	"fmt"
	"strings"
)

type person struct {
	name string
	age  uint
}
type personList []person

// supportin function for code reusability
func (list personList) findPersonByName(name string) (int, *person) {
	for index := range list {
		if strings.EqualFold(list[index].name, name) {
			return index, &list[index]
		}
	}
	return -1, nil
}

// checking if the person is eligible for voting
func (person person) checkVoteEligibility() bool {
	return person.age >= 18
}

// fetching  person details provided user must know the name
func (list personList) personsIntroduction() error {
	var name string
	fmt.Print("Enter the person name to see the details\n")
	fmt.Scan(&name)
	_, person := list.findPersonByName(name)
	if person == nil {
		return errors.New("Person not found!!")
	}

	fmt.Printf("Hello!!  %s  you are  %d years old\n", person.name, person.age)
	if person.checkVoteEligibility() {
		fmt.Println("you are eligible to vote")
	} else {
		fmt.Println("you are not eligible to vote")
	}
	return nil
}

// updating the age of person and sending message if updated age is eligible for voting
func (list *personList) updateAge() error {
	var name string
	fmt.Print("Enter name to update age ")
	fmt.Scan(&name)

	index, person := (*list).findPersonByName(name)
	if person == nil {
		return errors.New("Person not found!!")
	}
	var newAge uint
	fmt.Print("Enter new age")
	fmt.Scan(&newAge)
	if newAge < 1 || newAge > 140 {
		return errors.New("Please Enter valid age")
	}
	previousAge := (*list)[index].age
	(*list)[index].age = newAge
	fmt.Println("Your age has been updated successfully!")
	if previousAge < 18 && newAge >= 18 {
		fmt.Println("You are now eligible for voting")
	}
	return nil
}

// accepting the valid choice from user
func getChoice() (int, error) {
	var choice uint
	fmt.Println("Enter the choice")
	fmt.Println("1:person details || 2:update age || 3:exit")
	fmt.Scan(&choice)
	if choice < 1 || choice > 3 {
		return -1, errors.New("invalid selection")
	}
	return int(choice), nil
}

func main() {

	people := personList{
		{name: "Advait", age: 28},
		{name: "Atharv", age: 18},
		{name: "Aaryan", age: 11},
		{name: "Danish", age: 67},
		{name: "pratik", age: 18},
	}

	for {
		choice, err := getChoice()
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			if err := people.personsIntroduction(); err != nil {
				fmt.Println(err)
			}
		case 2:
			if err := people.updateAge(); err != nil {
				fmt.Println(err)
			}
		case 3:
			fmt.Print("Thank you!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
