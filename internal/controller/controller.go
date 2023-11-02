package controller

import (
	"HeadHunter/internal/models"
	"HeadHunter/internal/providers/db"
	"HeadHunter/pkg/employmentService"
	"fmt"
)

type Controller struct {
	db *db.DataBase
}

func NewController(base *db.DataBase) (*Controller, error) {

	return &Controller{db: base}, nil

}

func (c *Controller) AddEmployee(request *employmentService.EmployeeRequest) (string, error) {
	// adding employee and his skills
	err := c.db.AddEmployee(models.ConvertEmployeeRequestToModelsEmployee(request))
	if err != nil {
		return "unable to add the employee", err
	}

	return "employee was added", nil

}

func (c *Controller) AddCompanyByName(request *employmentService.CompanyRequest) (string, error) {
	// adding company by its name
	err := c.db.AddCompanyByName(request.Name)

	if err != nil {
		return "unable to add the company", err
	}

	return "company was added", nil
}

func (c *Controller) AddDepartmentToCompany(request *employmentService.DepartmentRequest) (string, error) {
	// adding department to company
	err := c.db.AddDepartmentToCompany(models.ConvertDepartmentRequestToModelsDepartment(request))
	if err != nil {
		return "unable to add the department", err
	}

	return "department was added", nil
}

func (c *Controller) EmployPerson(request *employmentService.EmploymentRequest) (string, error) {
	// employing person
	err := c.db.EmployPersonByEmail(request.Email, request.CompanyName, request.DepartmentName)
	if err != nil {
		return fmt.Sprintf("unable to employ person with email %s to dep %s", request.Email, request.DepartmentName), err
	}

	return "request was sent", nil

}
