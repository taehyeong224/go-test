package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	database = "test"
	user     = "root"
	password = "1234"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func Init() {
	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Variables for printing column data when scanned.
	var (
		id       int
		username     string
	)
	// Read some data from the table.
	rows, err := db.Query("SELECT id, username from auth_user;")
	checkError(err)
	defer rows.Close()
	fmt.Println("Reading data:")

	for rows.Next() {
		err := rows.Scan(&id, &username)
		checkError(err)
		fmt.Printf("Data row = (%d, %s)\n", id, username)
	}
	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")
}
