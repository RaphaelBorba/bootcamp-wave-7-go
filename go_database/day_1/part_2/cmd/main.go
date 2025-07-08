package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &application.ConfigDefault{
		Database: mysql.Config{
			User:      "root",
			Passwd:    "rootpass",
			Net:       "tcp",
			Addr:      "localhost:3306",
			DBName:    "storage_api_db",
			ParseTime: true,
		},
		Address: "localhost:8080",
	}
	app := application.NewDefault(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
