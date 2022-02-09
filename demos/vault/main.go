package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Log struct {
	ID         int64          `json:"log_id"`
	Type       string         `json:"log_type"`
	UserID     string         `json:"log_user_id"`
	UserIP     string         `json:"log_user_ip"`
	Time       sql.NullInt64  `json:"log_time"`
	Message    sql.NullString `json:"log_message"`
	ForeignKey sql.NullInt64  `json:"log_foreign_key"`
	History    sql.NullString `json:"log_history"`
}

func main() {
	fmt.Printf("Fakku Vault Demo\n")

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	address := os.Getenv("MYSQL_ADDRESS")
	database := os.Getenv("MYSQL_DATABASE")

	sqlConfig := mysql.NewConfig()
	sqlConfig.User = username
	sqlConfig.Passwd = password
	sqlConfig.Net = "tcp"
	sqlConfig.Addr = address
	sqlConfig.DBName = database
	sqlConfig.AllowOldPasswords = true
	sqlConfig.Timeout = 10 * time.Second
	sqlConfig.Params = map[string]string{
		"tls": "skip-verify",
	}

	connectionString := sqlConfig.FormatDSN()
	fmt.Printf("Connecting to MySQL database: %s\n",
		strings.Replace(connectionString,
			password,
			strings.Repeat("*", len(password)),
			-1,
		),
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, queryErr := db.Query("SELECT * FROM \"logs\"")
	if queryErr != nil {
		panic(queryErr.Error())
	}
	defer results.Close()
	fmt.Printf("Successfully connected to the Database at %s\n", sqlConfig.Addr)
	for results.Next() {
		var logEntry Log
		scanErr := results.Scan(
			&logEntry.ID,
			&logEntry.Type,
			&logEntry.UserID,
			&logEntry.UserIP,
			&logEntry.Time,
			&logEntry.Message,
			&logEntry.ForeignKey,
			&logEntry.History,
		)
		if scanErr != nil {
			fmt.Printf("Failed to scan results %v", scanErr)
			continue
		}
		fmt.Printf("LogEntry: %+v", logEntry)
	}
}
