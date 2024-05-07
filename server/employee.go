package server

import (
	"gocrud/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var emp storage.Employee

	if c.ShouldBind(&emp) == nil {
		log.Println("emp.Sender", emp.Name)
		id, err := emp.SaveEmployee()
		if err != nil {
			c.JSON(500, gin.H{"success": false, "emp": err.Error()})
			return
		}
		c.JSON(200, gin.H{"success": true, "ID": id})
		return
	}
}

func GetEmployeeByID(c *gin.Context) {
	var emp storage.Employee
	var data []storage.Employee

	if c.Bind(&emp) == nil {
		data, _ = emp.FetchEmployees(emp.Id)
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func DeleteEmployeeByID(c *gin.Context) {

	var emp storage.Employee
	var data bool

	if c.Bind(&emp) == nil {
		log.Println(" DeleteEmployeeByID emp.Sender", emp.Id)
		data, _ = emp.DeleteEmployee(emp.Id)

	}
	// emp.GetEmployees()

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}
