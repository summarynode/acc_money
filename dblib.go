package main

import "database/sql"

func mysqlConnect() *sql.DB  {
	// mysql connection
	db, err := sql.Open("mysql", "root:1softwise@tcp(erpyjun2.cafe24.com:3306)/day_data")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	return db
}


