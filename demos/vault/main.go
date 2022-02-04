package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hashicorp/vault/api"
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

	vaultServerAddress := os.Getenv("VAULT_SERVER_ADDRESS")
	vaultClient, vaultClientErr := api.NewClient(&api.Config{
		Address: vaultServerAddress,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})
	if vaultClientErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Vault client for address %s. %v\n", vaultServerAddress, vaultClientErr)
		os.Exit(2)
	}
	mysqlVaultSecretName := os.Getenv("MYSQL_VAULT_SECRET_NAME")

	mysqlCreds, vaultRespErr := vaultClient.Logical().Read(mysqlVaultSecretName)
	if vaultRespErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to get credential(%s) from Vault. %v\n", mysqlVaultSecretName, vaultRespErr)
		os.Exit(1)
	}

	fmt.Printf("Vault DB Creds: %#v\n", mysqlCreds.Data)
	fmt.Println("quitting...")
	os.Exit(0)

	sqlConfig := mysql.NewConfig()
	sqlConfig.User = os.Getenv("MYSQL_USERNAME")
	sqlConfig.Passwd = os.Getenv("MYSQL_PASSWORD")
	sqlConfig.Net = "tcp"
	sqlConfig.Addr = os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT")
	sqlConfig.DBName = os.Getenv("MYSQL_DATABASE")
	sqlConfig.AllowOldPasswords = true
	sqlConfig.Timeout = 10 * time.Second
	sqlConfig.Params = map[string]string{
		"tls": "skip-verify",
	}

	connectionString := sqlConfig.FormatDSN()
	fmt.Printf("Connecting to MySQL database: %s\n", connectionString)

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
		scanErr := results.Scan(&logEntry.ID, &logEntry.Type, &logEntry.UserID, &logEntry.UserIP, &logEntry.Time, &logEntry.Message, &logEntry.ForeignKey, &logEntry.History)
		if scanErr != nil {
			fmt.Printf("Failed to scan results %v", scanErr)
			continue
		}
		fmt.Printf("LogEntry: %+v", logEntry)
	}
}
