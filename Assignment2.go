package main

import (
	"errors"
	"fmt"
)

type employees struct {
	id           uint
	employeeName string
	age          uint
	salary       float64
}
type department struct {
	departmentName string
	employees      []employees
}

var currentId uint = 1

// getting index from helper function ensuring reusability
func (dept department) findIndex(id uint) int {
	for i, emp := range dept.employees {
		if emp.id == id {
			return i
		}
	}
	return -1
}

// fetching all employees
func (dept department) employeeList() {
	for _, emp := range dept.employees {
		fmt.Printf("ID: %d,  Name: %s,  Age: %d,  Salary: %.2f,  Department:%s\n",
			emp.id, emp.employeeName, emp.age, emp.salary, dept.departmentName)
	}
}

// adding employee by taking details from user with error handeling
func (dept *department) addEmployee() {
	fmt.Print("please enter the employee details\n")
	var name string
	var age uint
	var salary float64
	fmt.Println("enter the employee's name")
	fmt.Scan(&name)

	fmt.Println("enter the employee's age")
	if _, err := fmt.Scan(&age); err != nil {
		fmt.Println("please Enter valid age")
		return
	}

	fmt.Println("enter the employee's salary")
	if _, err := fmt.Scan(&salary); err != nil {
		fmt.Println("please Enter valid salary")
		return
	}

	employee := employees{
		id:           currentId,
		employeeName: name,
		age:          age,
		salary:       salary,
	}
	dept.employees = append(dept.employees, employee)
	currentId++
	fmt.Println("Employee added!")

}

// removing empoyee based on employee id
func (dept *department) removeEmployee() error {
	fmt.Println("Enter the employee id to delete")
	var empId uint
	fmt.Scan(&empId)
	index := dept.findIndex(empId)
	if index == -1 {
		return errors.New("Invalid employee ID.")

	}
	dept.employees = append(dept.employees[:index], dept.employees[index+1:]...)

	fmt.Println("Employee removed!")
	return nil

}

// updating salary based on id and raise amount
func (dept *department) giveRaise() error {
	var raiseAmount float64
	var empId uint
	fmt.Println("Enter employee ID")
	fmt.Scan(&empId)

	index := dept.findIndex(empId)
	if index == -1 {
		return errors.New("Invalid employee ID.")
	}
	fmt.Println("Enter raise amount")
	fmt.Scan(&raiseAmount)

	dept.employees[index].salary += raiseAmount
	fmt.Println("Salary has been updated!")
	return nil
}

// calculating average salary of department
func (dept department) calculateAverageSalary() {
	count := len(dept.employees)
	var sum float64
	for _, emp := range dept.employees {
		sum += emp.salary
	}

	fmt.Printf("Average Salary of Department: %.2f\n", sum/float64(count))

}

// accepting choice from user
func getUserChoice() (int, error) {
	var choice uint
	fmt.Println("\nEnter choice to perform operation ")
	fmt.Println("1:Employee List || 2:Add Employee || 3:Remove Employee || 4:Give Raise || 5:Average Salary || 0:Exit")
	fmt.Scan(&choice)
	if choice > 5 {
		return -1, errors.New("Please enter a valid choice")

	}
	return int(choice), nil
}

func main() {
	dept := department{
		departmentName: "Computer Science",
		employees: []employees{
			{id: currentId, employeeName: "Atharv", age: 22, salary: 40000},
			employees{id: currentId + 1, employeeName: "pranav", age: 21, salary: 20000},
			employees{id: currentId + 2, employeeName: "Aaryan", age: 17, salary: 10000},
		},
	}
	//currentId += 3

	for {
		choice, err := getUserChoice()
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch choice {
		case 1:
			dept.employeeList()

		case 2:
			dept.addEmployee()
		case 3:
			if err := dept.removeEmployee(); err != nil {
				fmt.Println(err)
			}
		case 4:
			if err := dept.giveRaise(); err != nil {
				fmt.Println(err)
			}
		case 5:
			dept.calculateAverageSalary()
		case 0:
			fmt.Println("Thank you!!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}
