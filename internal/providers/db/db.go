package db

import (
	"HeadHunter/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DataBase struct {
	db *sql.DB
}

func NewDB(host, user, password, nameOfDB, driverName string, port int) (*DataBase, error) {
	// constructor for DataBase struct
	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, nameOfDB)
	if result, err := sql.Open(driverName, params); err != nil {
		return nil, err
	} else {
		if err = result.Ping(); err != nil {
			return nil, err
		}
		dataBase := DataBase{result}
		_, err := dataBase.db.Query("DELETE FROM public.workers;")
		if err != nil {
			return nil, err
		}
		return &dataBase, nil
	}

}

func (db *DataBase) AddCompanyByName(name string) error {
	// adding company by name
	queryString := fmt.Sprintf("insert into company(company_name)\nvalues('%s')", name)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}
	return nil

}

func (db *DataBase) AddEmployee(employee models.Employee) error {
	// adding employee and his skills
	queryString := fmt.Sprintf("insert into employee(emp_name)\nvalues('%s')", employee.Name)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}
	for skill, _ := range employee.Skills {
		queryString := fmt.Sprintf("INSERT INTO skill (skill, emp_id) \nSELECT '%s', emp_id FROM employee WHERE emp_name = '%s';", skill, employee.Name)
		_, err = db.db.Query(queryString)
	}
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) AddDepartmentToCompany(companyName string, department models.Department) error {
	// adding department to existing company
	queryString := fmt.Sprintf("insert into department\n(dep_name, company_id)\nselect '%s', company_id from company where company_name = '%s'", department.Name, companyName)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}
	for skill, _ := range department.RequiredSkills {
		queryString = fmt.Sprintf("insert into required(skill, dep_id)\nselect '%s', dep_id from department where dep_name = '%s'", skill, department.Name)
		_, err = db.db.Query(queryString)
	}
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) EmployPersonByName(name string, companyName string, departmentName string) error {
	// employing person to the department of the company
	queryString := fmt.Sprintf(
		"update employee "+
			"set dep_id = subquery.dep_id"+
			" from (select dep_id  from department where dep_name = '%s' and company_id = (select company_id from company where company_name = '%s')) as subquery "+
			"where emp_name = '%s' and employee.dep_id is null "+
			"and (select count(*) from required join skill using (skill) where required.dep_id = "+
			" (select dep_id from department where dep_name = '%s' "+
			"  and company_id = (select company_id from company where company_name = '%s')) "+
			" and skill.emp_id = (select emp_id from employee where emp_name = '%s')) = "+
			" (select count(*) from required where dep_id = (select dep_id from department where dep_name = '%s' "+
			" and company_id = (select company_id from company where company_name = '%s')));", departmentName, companyName, name, departmentName, companyName, name, departmentName, companyName)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}
