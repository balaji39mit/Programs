package main

import "fmt"

var (
	n                int
	id, supervisorid int64
	name, position   string
	salary, bonus    float64
	rating           uint8
)

func createEmployee(name string, id int64, salary float64, rating uint8, position string, supervisorId int64) {
	employee := &Employee{
		Name:         name,
		Id:           id,
		Salary:       salary,
		Rating:       Rating(rating),
		Position:     Position(position),
		SupervisorId: supervisorid,
	}
	employeeManager.InsertEmployee(employee)
}

func main() {
	fmt.Println("Enter the number of employees : ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println(fmt.Sprintf("Enter the details of the employee #%v in the following order : Name, Id, Salary, Rating, Position and SupervisorId..", i+1))
		fmt.Scanln(&name, &id, &salary, &rating, &position, &supervisorid)
		createEmployee(name, id, salary, rating, position, supervisorid)
	}
	employeeManager.DisplayEmployee()
	//processHierarchicalDisplay()
	processBonus()
	printPerformers()
}

func printPerformers() {
	fmt.Println("Please enter the number of best employees you want : ")
	fmt.Scanln(&n)
	fmt.Println("Printing best permance of employees at DE.")
	employeeManager.TopPerformers(n, DeliveryExecutives)
}

func processHierarchicalDisplay() {
	fmt.Println("Enter the employee id from which hierarachy have to be printed : ")
	fmt.Scanln(&id)
	employeeManager.DisplayHierarchy(id)
}

func processBonus() {
	fmt.Println("Enter the bonus amount :")
	fmt.Scanln(&bonus)
	employeeManager.ProcessBonus(bonus)
}
