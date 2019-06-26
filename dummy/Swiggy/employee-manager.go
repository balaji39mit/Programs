package main

import (
	"fmt"
)

var employeeManager IEmployeeManager

func init() {
	employeeManager = &Management{}
}

func (manager *Management) FindEmployee(Id int64) *Employee {
	if manager.Root == nil {
		return nil
	}
	return manager.Root.Find(Id)
}

func (manager *Management) InsertEmployee(e *Employee) {
	if manager.Root == nil {
		fmt.Println("Creating the first employee of the management.")
		manager.Root = e
		return
	}
	//Check for the supervisor node.
	supervisor := manager.FindEmployee(e.SupervisorId)
	if supervisor != nil {
		supervisor.Children = append(supervisor.Children, e)
	} else {
		fmt.Println("Given employee is neither the first or child employee, Please check the inputs.")
	}
}

func (manager *Management) DisplayEmployee() {
	fmt.Println("Displaying the employees")
	manager.Root.Print(0)
}

func (manager *Management) DisplayHierarchy(Id int64) {
	employee := manager.Root.Find(Id)
	if employee != nil {
		employee.Print(0)
	} else {
		fmt.Println("Sorry, given employee is not present.")
	}
}

func (manager *Management) ProcessBonus(bonus float64) {
	//TODO: Find the employee based on the city and call splitBonus with that employee.
	if manager.Root == nil {
		return
	}
	manager.Root.SplitBonus(bonus)
}

/*func (manager *Management) FindLowestLevel() Position {
	return manager.Root.LowestLevel()
}*/

func (manager *Management) TopPerformers(number int, position Position) {
	performers := make(map[float64]*Employee, number)
	manager.Root.AnalyzePerformance(&performers, number, position)
	fmt.Println(fmt.Sprintf("The top Employees in %v are :", string(position)))
	for _, val := range performers {
		fmt.Sprintf("Employee's : %v", val.Name)
	}
}
