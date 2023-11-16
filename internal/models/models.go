package models

// structures for all project

type Employee struct {
	Name   string
	Email  string
	Skills []string
}
type EmployeeInfo struct {
	Name           string
	Email          string
	CompanyName    string
	DepartmentName string
	Status         string
	Salary         int
}
type Department struct {
	Name                       string
	RequiredSkills             []string
	EmployeesLimit             int
	PromotionIntervalInMinutes int
}
