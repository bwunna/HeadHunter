package models

import "HeadHunter/pkg/employmentService"

func ConvertEmployeeRequestToModelsEmployee(request *employmentService.EmployeeRequest) Employee {
	skills := make(map[string]interface{})
	for skill := range request.Skills {
		var b interface{}
		skills[skill] = b
	}
	return Employee{Name: request.Name, Skills: skills}
}

func ConvertDepartmentRequestToModelsDepartment(request *employmentService.DepartmentRequest) (string, Department) {
	skills := make(map[string]interface{})

	for skill := range request.RequiredSkills {
		var b interface{}
		skills[skill] = b
	}

	return request.CompanyName, Department{request.DepName, skills}
}
