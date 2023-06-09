package routes
import (
    "github.com/balram1312/go-gin-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/go-pg/pg/v9"
    "log"
    orm "github.com/go-pg/pg/v9/orm"
    "strconv"
	"os"
)


func EmployeeRouter(router *gin.Engine){

    router.GET("/employees", GetAllEmployees)
	router.POST("/employee", CreateEmployee)
    router.DELETE("/employee/:employeeID", DeleteEmployee)

}

func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "tiger",
		Addr: "localhost:5432",
		Database: "employeedb",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	CreateEmployeeTable(db)
	InitiateDB(db)
	return db
}

func CreateEmployeeTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Employee{}, opts)
	if createError != nil {
		log.Printf("Error while creating employee table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Employee table created")
	return nil
}

var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}



func CreateEmployee(c *gin.Context) {
    var employee models.Employee
	c.BindJSON(&employee)
    //ID := employee.ID
    name := employee.Name      
    leavetype := employee.Leavetype
    fromdate := employee.Fromdate
    todate :=  employee.Todate
    teamname := employee.Teamname   
    file := employee.File
    reporter :=  employee.Reporter  

	insertError := dbConnect.Insert(&models.Employee{
       
        Name: name,     
        Leavetype: leavetype, 
        Fromdate: fromdate,
        Todate: todate,
        Teamname: teamname,
        File: file,
        Reporter: reporter, 
	})
	if insertError != nil {
		log.Printf("Error while inserting new employee into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Employee created Successfully",
	})
	return
}


func GetAllEmployees(c *gin.Context) {
	var employees []models.Employee
	err := dbConnect.Model(&employees).Select()

	if err != nil {
		log.Printf("Error while getting all employees, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Employees",
		"data": employees,
	})
	return
}



func DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("employeeID")
	employeeIDInt, err := strconv.Atoi(employeeID)
	if err != nil {
		log.Printf("Invalid employee ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid employee ID",
		})
		return
	}

	employee := &models.Employee{ID: employeeIDInt}

	err = dbConnect.Delete(employee)
	if err != nil {
		log.Printf("Error while deleting an employee, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Employee deleted successfully",
	})
}