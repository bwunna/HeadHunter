package controller

import (
	"HeadHunter/internal/models"
	"HeadHunter/internal/providers/db"
	"HeadHunter/pkg/employmentService"
)

type Controller struct {
	db *db.DataBase
}

// controller between api and db

func NewController(base *db.DataBase) *Controller {
	return &Controller{
		db: base,
	}
}

// sending request to db to add employee and his skills

func (c *Controller) AddEmployee(request *employmentService.EmployeeRequest) error {
	// adding employee and his skills
	err := c.db.AddEmployee(models.ConvertEmployeeRequestToModelsEmployee(request))
	if err != nil {
		return err
	}

	return nil
}

// sending request to db to add company by its name

func (c *Controller) AddCompanyByName(companyName string) error {
	// adding company by its name
	err := c.db.AddCompanyByName(companyName)
	if err != nil {
		return err
	}

	return nil
}

// sending request to db to update the salary of employee

func (c *Controller) UpdateSalaryByEmail(email string) error {
	err := c.db.UpdateSalaryByEmail(email)
	return err
}

// sending request to db to add company

func (c *Controller) AddDepartmentToCompany(request *employmentService.DepartmentRequest) error {
	err := c.db.AddDepartmentToCompany(models.ConvertDepartmentRequestToModelsDepartment(request))
	if err != nil {
		return err
	}

	return nil
}

// sending request to db to employ person

func (c *Controller) EmployPerson(email string, departmentName string, companyName string) error {
	// employing person
	err := c.db.EmployPersonByEmail(email, departmentName, companyName)
	return err

}

// sending request to db to get employee info

func (c *Controller) GetEmployeeInfoByEmail(email string) (*employmentService.EmployeeInfo, error) {
	info, err := c.db.GetEmployeeInfoByEmail(email)
	if err != nil {
		return nil, err
	}

	return models.ConvertToGRPCEmployeeInfo(info), nil
}
