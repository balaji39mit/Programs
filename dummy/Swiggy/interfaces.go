package main

type IEmployeeManager interface {
	FindEmployee(int64) *Employee
	InsertEmployee(*Employee)
	DisplayEmployee()
	DisplayHierarchy(int64)
	ProcessBonus(float64)
	TopPerformers(int, Position)
}

type IEmployee interface {
	Find(int64) *Employee
	Print(int)
	SplitBonus(float64)
	AnalyzePerformance(map[float64]*Employee, int, Position)
}
