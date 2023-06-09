package config

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/balram1312/go-gin-api/models"
)

var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
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
	log.Printf("Todo table created")
	return nil
}
// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "postgres",
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

