package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Table    string `json:"table"`
}

func LoadConfig(configPath string) (*DBConfig, error) {
	filePath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config DBConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func Connect(config *DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=false",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	fmt.Printf("Attempting to connect to database: %s\n", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error creating database connection: %v\n", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Database ping failed: %v\n", err)
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
