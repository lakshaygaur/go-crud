package storage

import (
	"fmt"
	"log"
)

type Employee struct {
	Id       int     `form:"id" json:"id"`
	Name     string  `form:"name" json:"name"`
	Position string  `form:"position" json:"position"`
	Salary   float32 `form:"salary" json:"salary"`
}

func (emp Employee) SaveEmployee() (int64, error) {
	result, err := DbHandler.db.Exec("INSERT INTO employees ( name, position, salary) VALUES ( ?, ?, ?)", emp.Name, emp.Position, emp.Salary)
	if err != nil {
		log.Println("failed creating employee", err)
		return 0, err
	}
	fmt.Println("In between result ")
	id, err := result.LastInsertId()
	fmt.Println("ID of inserted chat", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (emp Employee) FetchEmployees(empId int) ([]Employee, error) {
	// query DB
	data, err := DbHandler.db.Query("select * from employees where id=?", empId)
	if err != nil {
		log.Fatal("Failed executing FetchEmployees query ", err)
		return nil, err
	}
	employees := []Employee{}
	for data.Next() {
		emp := Employee{}
		data.Scan(&emp.Id, &emp.Name, &emp.Position, &emp.Salary)
		employees = append(employees, emp)
	}
	return employees, nil
}

func (emp Employee) DeleteEmployee(empId int) (bool, error) {
	_, err := DbHandler.db.Exec("delete from employees where id=?", empId)
	if err != nil {
		log.Fatal("Failed executing DeleteEmployee query ", err)
		return false, err
	}
	return true, nil
}
