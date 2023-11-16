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

// converts slice of keys into

func KeysInString(keys []string) string {
	var answer string
	for _, value := range keys {
		answer += fmt.Sprintf("'%s',", value)
	}

	if len(answer) > 1 {
		answer = answer[:len(answer)-1]
	}
	return answer
}

// constructor for db

func NewDB(host, user, password, nameOfDB, driverName string, port int) (*DataBase, error) {
	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, nameOfDB)

	result, err := sql.Open(driverName, params)
	if err != nil {
		return nil, err
	}

	err = result.Ping()
	if err != nil {
		return nil, err
	}

	dataBase := DataBase{result}

	return &dataBase, nil

}

// adding company by its name

func (base *DataBase) AddCompanyByName(name string) error {
	queryString := fmt.Sprintf(`insert into company(company_name) values('%s')`, name)
	_, err := base.db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

// adding employee and his skills

func (base *DataBase) AddEmployee(employee models.Employee) error {
	queryString := fmt.Sprintf(`insert into employee(emp_name, skills, email) values('%s', array[%s], '%s')`, employee.Name, KeysInString(employee.Skills), employee.Email)
	_, err := base.db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

// adding department to existing company

func (base *DataBase) AddDepartmentToCompany(companyName string, department models.Department) error {
	queryString := fmt.Sprintf(`insert into department (dep_name, company_id, required_skills, emp_limit, prom_interval) select '%s', 
                                                                      company_id, array[%s], %d, '%d minutes' from company where company_name = '%s'`,
		department.Name, KeysInString(department.RequiredSkills), department.EmployeesLimit, department.PromotionIntervalInMinutes, companyName)
	_, err := base.db.Query(queryString)
	if err != nil {
		return err
	}

	return nil
}

// updating salary if employee has worked enough

func (base *DataBase) UpdateSalaryByEmail(email string) error {
	queryString := fmt.Sprintf(`update employee set current_salary = raise_salary(current_salary), status = next_status(status)
	where email = '%s' and now() > salary_update_time + (select prom_interval from 
			department where dep_id = (select dep_id from employee where 
				email = '%s')) `, email, email)
	_, err := base.db.Query(queryString)
	return err

}

// employing person if his skills matching department skills

func (base *DataBase) EmployPersonByEmail(email string, departmentName string, companyName string) error {
	queryString := fmt.Sprintf(
		`do $$
	declare 
	var_dep_id integer;
	var_company_id integer;
	var_emp_id integer;
	emp_amount integer;
	var_emp_limit integer;
	begin
	select company_id into var_company_id from company where company_name = '%s';
	select dep_id into var_dep_id from department where dep_name = '%s' and company_id = var_company_id;
	select emp_id into var_emp_id from employee where email = '%s';
	select count(*) into emp_amount from employee where dep_id = var_dep_id;
	select emp_limit into var_emp_limit from department where dep_id = var_dep_id;
	update employee set dep_id = var_dep_id, current_salary = make_salary(current_salary), 
	avg_salary = current_salary, status = make_rank(status), salary_update_time = now()
	where (dep_id is null or dep_id != var_dep_id) and emp_amount < var_emp_limit and emp_id = var_emp_id and 
	(select (select count(*)	from ((select unnest(skills)
	from employee where emp_id = var_emp_id)
	join (select unnest(required_skills)
	from department
	where dep_id = var_dep_id)
	using (unnest))) = (select count(*) from (select unnest(required_skills)
	from department where dep_id = var_dep_id)));
	end $$;`,
		companyName, departmentName, email)
	_, err := base.db.Query(queryString)
	if err != nil {
		return err
	}

	return nil
}

// getting info about employee name, email, salary, status and place of work

func (base *DataBase) GetEmployeeInfoByEmail(email string) (*models.EmployeeInfo, error) {
	queryString := fmt.Sprintf(
		`select emp_name, email, (select dep_name from department where dep_id = employee.dep_id), 
       			(select company_name
				from company where company_id = 
				(select company_id from department where dep_id = employee.dep_id)), status, current_salary
				from employee 
				where email = '%s';`, email)
	row, err := base.db.Query(queryString)
	if err != nil {
		return nil, err
	}

	info := &models.EmployeeInfo{}
	for row.Next() {
		err = row.Scan(&info.Name, &info.Email, &info.DepartmentName, &info.CompanyName, &info.Status, &info.Salary)
		if err != nil {
			if info.Name != "" || info.Email != "" {
				return info, nil
			}

			return nil, err
		}
	}
	if info.Email == "" || info.Name == "" {
		return nil, fmt.Errorf("employee not found")
	}
	return info, nil

}
