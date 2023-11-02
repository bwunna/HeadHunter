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

func KeysInString(keys []string) string {
	var answer string
	for _, value := range keys {
		answer += "'" + value + "', "
	}
	if len(answer) > 2 {
		answer = answer[:len(answer)-2]
	}
	return answer
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
	queryString := fmt.Sprintf("insert into employee(emp_name, skills, email) values('%s', array[%s], '%s')", employee.Name, KeysInString(employee.Skills), employee.Email)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}

	return nil
}

func (db *DataBase) AddDepartmentToCompany(companyName string, department models.Department) error {
	// adding department to existing company
	queryString := fmt.Sprintf("insert into department (dep_name, company_id, required_skills) select '%s', company_id, array[%s] from company where company_name = '%s'", department.Name, KeysInString(department.RequiredSkills), companyName)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}

	return nil
}

func (db *DataBase) EmployPersonByEmail(email string, companyName string, departmentName string) error {
	// employing person to the department of the company

	queryString := fmt.Sprintf(
		"update employee set dep_id = subquery.dep_id from (select dep_id from department where dep_name = '%s'"+
			" and company_id = (select company_id from company where"+
			" company_name = '%s')) as subquery where email = '%s' and (select (select count(*) "+
			"from ((select unnest(skills)"+
			" from employee where email = '%s')"+
			" join (select unnest(required_skills)"+
			" from department"+
			" where dep_name = '%s' and company_id = "+
			" (select company_id from company where company_name = '%s')) "+
			" using (unnest))) = (select count(*) from (select unnest(required_skills)"+
			" from department where dep_name = '%s' and company_id = (select company_id from company where company_name = '%s'))))", departmentName, companyName, email, email, departmentName, companyName, departmentName, companyName)
	_, err := db.db.Query(queryString)
	if err != nil {
		return err
	}

	return nil
}
