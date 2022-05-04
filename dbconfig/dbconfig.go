package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConnect()  {
	// Database connection to DB and connect the db
	cfg := mysql.Config{
		User: "root",
		Passwd: "",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "learn3012",
		AllowNativePasswords: true,
	}

	// Get the db handler
	var err error

	DB, err = sql.Open("mysql",cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Print("Database connected\n")

    defer DB.Close()
}

