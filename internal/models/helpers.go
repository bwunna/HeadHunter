package models

import (
	"HeadHunter/pkg/employmentService"
)

//

func ConvertEmployeeRequestToModelsEmployee(request *employmentService.EmployeeRequest) Employee {

	return Employee{
		Name:   request.Name,
		Skills: request.Skills,
		Email:  request.Email,
	}
}

func ConvertToGRPCEmployeeInfo(c *EmployeeInfo) *employmentService.EmployeeInfo {
	return &employmentService.EmployeeInfo{
		Name:           c.Name,
		DepartmentName: c.DepartmentName,
		CompanyName:    c.CompanyName,
		Salary:         int32(c.Salary),
		Status:         c.Status,
		Email:          c.Email,
	}
}

func ConvertDepartmentRequestToModelsDepartment(request *employmentService.DepartmentRequest) (string, Department) {

	return request.CompanyName, Department{
		Name:                       request.DepName,
		RequiredSkills:             request.RequiredSkills,
		EmployeesLimit:             int(request.EmployeeLimit),
		PromotionIntervalInMinutes: int(request.PromotionIntervalInMinutes),
	}
}
