package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Employees-Management/initializers"
	"github.com/waltertaya/Employees-Management/models"
)

func PostEmployees(c *gin.Context) {
	// Fetch json data send
	var body struct {
		FirstName string
		LastName string
		Email string
		ContactNumber string
		Age int
		Dob string
		Salary int
		Address string
	}
	c.Bind(&body)

	// create employee
	employee := models.Employees{
		FirstName: body.FirstName,
		LastName: body.LastName,
		Email: body.Email,
		ContactNumber: body.ContactNumber,
		Age: body.Age,
		Dob: body.Dob,
		Salary: body.Salary,
		Address: body.Address,
	}

	result := initializers.DB.Create(&employee)

	// Return
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"employee": employee,
	})
}

func GetEmployees(c *gin.Context) {
	// Fetch all employees
	var employees []models.Employees
	initializers.DB.Find(&employees)

	// Return
	c.JSON(200, gin.H{
		"employees": employees,
	})
}

func GetEmployee(c *gin.Context) {
	// Fetch employee by id
	var employee models.Employees
	initializers.DB.First(&employee, c.Param("id"))

	// Return
	c.JSON(200, gin.H{
		"employee": employee,
	})
}

func UpdateEmployee(c *gin.Context) {
	// Fetch json data send
	var body struct {
		FirstName string
		LastName string
		Email string
		ContactNumber string
		Age int
		Dob string
		Salary int
		Address string
	}
	c.Bind(&body)

	// Fetch employee by id
	var employee models.Employees
	initializers.DB.First(&employee, c.Param("id"))

	// Update employee
	employee.FirstName = body.FirstName
	employee.LastName = body.LastName
	employee.Email = body.Email
	employee.ContactNumber = body.ContactNumber
	employee.Age = body.Age
	employee.Dob = body.Dob
	employee.Salary = body.Salary
	employee.Address = body.Address

	initializers.DB.Save(&employee)

	// Return
	c.JSON(200, gin.H{
		"employee": employee,
	})
}

func DeleteEmployee(c *gin.Context) {
	// Fetch employee by id
	var employee models.Employees
	initializers.DB.First(&employee, c.Param("id"))

	// Delete employee
	initializers.DB.Delete(&employee)

	// Return
	c.JSON(200, gin.H{
		"employee": employee,
	})
}
