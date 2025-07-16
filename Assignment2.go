package main

import (
	"fmt"
)

type Employee struct {
	employeeId     int
	employeeName   string
	employeeAge    int
	employeeSalary float64
}
type Department struct {
	departmentName string
	employees      map[int]Employee
}

// storing employees in a map
var employees = make(map[int]Employee)

// maintaining current id for adding employee in future
var currentId = 1

// storing values in employee map
var _ = func() bool {
	employees[currentId] = Employee{employeeId: currentId, employeeName: "Atharv", employeeAge: 22, employeeSalary: 40000}
	currentId++
	employees[currentId] = Employee{employeeId: currentId, employeeName: "pranav", employeeAge: 21, employeeSalary: 20000}
	currentId++
	employees[currentId] = Employee{employeeId: currentId, employeeName: "Aaryan ", employeeAge: 17, employeeSalary: 10000}
	currentId++
	return true
}()

// declaring department variable and initializing in init function to get loaded before main method
var dept Department

func init() {
	dept = Department{
		departmentName: "Computer Science",
		employees: map[int]Employee{
			1: employees[1],
			2: employees[2],
			3: employees[3],
		},
	}
}

// fetching all employees
func (dept *Department) employeeList() {
	for _, emp := range dept.employees {
		fmt.Printf("ID: %d,  Name: %s,  Age: %d,  Salary: %.2f,  Department:%s\n",
			emp.employeeId, emp.employeeName, emp.employeeAge, emp.employeeSalary, dept.departmentName)
	}
}

// adding employee by taking details from user
func (debt *Department) addEmployee() {
	fmt.Print("please enter the employee details\n")
	var employeeName string
	var employeeAge int
	var employeeSalary float64
	fmt.Println("enter the employee's name")
	fmt.Scan(&employeeName)

	fmt.Println("enter the employee's age")
	fmt.Scan(&employeeAge)

	fmt.Println("enter the employee's salary")
	fmt.Scan(&employeeSalary)

	employee := Employee{
		employeeId:     currentId,
		employeeName:   employeeName,
		employeeAge:    employeeAge,
		employeeSalary: employeeSalary,
	}
	debt.employees[currentId] = employee
	employees[currentId] = employee
	currentId++
	fmt.Println("Employee added!")

}

// removing empoyee based on employee id
func (dept *Department) removeEmployee() {
	fmt.Println("Enter the employee id to delete")
	var empId int
	fmt.Scan(&empId)
	if _, exists := employees[empId]; exists {
		delete(dept.employees, empId)
		delete(employees, empId)
		fmt.Println("Employee removed!")
	} else {
		fmt.Println("invalid employee id")
	}

}

// updating salary based on id and raise amount
func (dept *Department) giveRaise() {
	var raiseAmount float64
	var empId int

	fmt.Println("Enter the employee ID for upraisal")
	fmt.Scan(&empId)

	fmt.Println("Enter the upraisal amount")
	fmt.Scan(&raiseAmount)
	if emp, exists := employees[empId]; exists {
		emp.employeeSalary += raiseAmount
		dept.employees[empId] = emp
		employees[empId] = emp
		fmt.Println("Salary has been updated!")
	} else {
		fmt.Println("Invalid employee id entered")
	}
}

// calculating average salary of department
func (dept *Department) calculateAverageSalary() {
	count := len(dept.employees)
	sum := 0.0
	for _, emp := range dept.employees {
		sum += emp.employeeSalary
	}

	fmt.Printf("Average Salary of Department: %.2f\n", sum/float64(count))

}

func main() {

	for {
		var choice int
		fmt.Println("Enter choice to perform operation ")
		fmt.Println("1:Employee List || 2:Add Employee || 3:Remove Employee || 4:Give Raise || 5:Average Salary || 0:Main Menu")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			dept.employeeList()

		case 2:
			dept.addEmployee()
		case 3:
			dept.removeEmployee()
		case 4:
			dept.giveRaise()
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
