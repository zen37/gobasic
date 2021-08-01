package main

// https://zetcode.com/golang/mysql/

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//go_user has only SELECT privileges
	db, err := sql.Open("mysql", "go_user:<password>@tcp(127.0.0.1:3306)/ci_tutorial")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}
