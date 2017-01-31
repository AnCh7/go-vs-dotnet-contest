package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func loadEmployees() *Employee {

	rows, err := db.Query("select emp_no, first_name, last_name from employees where emp_no = 10415")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	emp := new(Employee)
	for rows.Next() {
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.LastName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(emp)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return emp
}
