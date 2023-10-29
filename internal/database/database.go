package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

const (
	sqlCreateTable = "CREATE TABLE IF NOT EXISTS users (id VARCHAR(250) PRIMARY KEY, name VARCHAR(100) NOT NULL, email VARCHAR(200) NOT NULL);"
	driver         = "mysql"
)

type Database interface {
	//GetConnection Get a database connection
	GetConnection() (*sql.DB, error)
	//Init Create database tables
	Init(db *sql.DB) error
}

type database struct {
}

func (r *database) GetConnection() (*sql.DB, error) {
	dsn := getDsn()
	print(dsn)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
func (r *database) Init(db *sql.DB) error {
	_, err := db.Exec(sqlCreateTable)
	if err != nil {
		return err
	}
	return nil
}

func getDsn() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbServer := os.Getenv("DB_SERVER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbServer, dbPort, dbName)
	return dsn
}

func NewDatabase() Database {
	result := &database{}
	return result
}
