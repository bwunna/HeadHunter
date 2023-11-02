package models

import (
	"HeadHunter/pkg/employmentService"
)

//

func ConvertEmployeeRequestToModelsEmployee(request *employmentService.EmployeeRequest) Employee {
	skills := make([]string, len(request.Skills))
	for index, skill := range request.Skills {
		skills[index] = skill
	}
	return Employee{Name: request.Name, Skills: skills, Email: request.Email}
}

func ConvertDepartmentRequestToModelsDepartment(request *employmentService.DepartmentRequest) (string, Department) {
	skills := make([]string, len(request.RequiredSkills))
	for index, skill := range request.RequiredSkills {
		skills[index] = skill
	}

	return request.CompanyName, Department{request.DepName, skills}
}
