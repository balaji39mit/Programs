package main

import "fmt"

type Employee struct {
	Id           int64
	Name         string
	City         string
	Salary       float64
	Rating       Rating
	Position     Position
	SupervisorId int64
	Children     []*Employee
	Bonus        float64
}

type Management struct {
	Root *Employee
}

func (e *Employee) SPrintf(format string, a ...interface{}) string {
	str := ""
	str += fmt.Sprintf("Id:%v", e.Id)
	str += fmt.Sprintf("Name:%v", e.Name)
	str += fmt.Sprintf("Salary:%v", e.Salary)
	str += fmt.Sprintf("Rating:%v", e.Rating)
	str += fmt.Sprintf("Position:%v", e.Position)
	str += fmt.Sprintf("Bonus:%v", e.Bonus)
	return str
}
