package DatabaseInitialization

import (
	yaml "TheGreatestFinanceManagerEver/Infrastructure/Config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func openDatabase(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func readSQLFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func applySqlFile(db *sql.DB, filePath string) error {
	sqlContent, err := readSQLFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlContent)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var port = yaml.Config.Database.Port
	var host = yaml.Config.Database.Host
	var dbname = yaml.Config.Database.DatabaseName
	var username = yaml.Config.Database.UserName
	var password = yaml.Config.Database.Password

	var connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)
	log.Print("connection string: " + connectionString)
	db, err := openDatabase(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	const sqlFilePath = "source/sql/000_start.sql"
	err = applySqlFile(db, sqlFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
