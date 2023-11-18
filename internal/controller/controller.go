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

func New(base *db.DataBase) *Controller {
	return &Controller{
		db: base,
	}
}

func (c *Controller) AddEmployee(request *employmentService.EmployeeRequest) error {
	// adding employee and his skills
	err := c.db.AddEmployee(models.ConvertEmployeeRequestToModelsEmployee(request))
	return err
}

func (c *Controller) AddCompanyByName(companyName string) error {
	// adding company by its name
	err := c.db.AddCompanyByName(companyName)
	return err
}

func (c *Controller) UpdateSalaryByEmail(email string) error {
	err := c.db.UpdateSalaryByEmail(email)
	return err
}

func (c *Controller) AddDepartmentToCompany(request *employmentService.DepartmentRequest) error {
	err := c.db.AddDepartmentToCompany(models.ConvertDepartmentRequestToModelsDepartment(request))
	return err
}

func (c *Controller) EmployPerson(email string, departmentName string, companyName string) error {
	// employing person
	err := c.db.EmployPersonByEmail(email, departmentName, companyName)
	return err

}

func (c *Controller) GetEmployeeInfoByEmail(email string) (*employmentService.EmployeeInfo, error) {
	info, err := c.db.GetEmployeeInfoByEmail(email)
	if err != nil {
		return nil, err
	}

	return models.ConvertToGRPCEmployeeInfo(info), nil
}
