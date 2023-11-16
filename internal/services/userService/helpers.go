package userService

import (
	"HeadHunter/pkg/employmentService"
	"HeadHunter/pkg/userService"
)

func ConvertToUserServiceEmployee(info *employmentService.EmployeeInfo) *userService.Employee {
	return &userService.Employee{
		Name:           info.Name,
		Email:          info.Email,
		Status:         info.Status,
		Salary:         int32(info.Salary),
		CompanyName:    info.CompanyName,
		DepartmentName: info.DepartmentName,
	}
}
