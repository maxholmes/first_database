package main

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	/*
	   variables required for connection string: connStr

	   user= (using default user for postgres database)
	   dbname= (using default database that comes with postgres)
	   password = (password used during initial setup)
	   host = (hostname or IP Address of server)
	   sslmode = (must be set to disabled unless using SSL)
	*/

	connStr := "user=postgres dbname=my_db password=*Password* host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	// insert a row
	sqlStatement := `INSERT INTO players (first_name, last_name, position) 
VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, "Julian", "Pellman", "Fullback")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}

}
