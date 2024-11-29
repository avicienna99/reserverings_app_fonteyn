package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig stores database connection settings
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Table    string `json:"table"`
}

// LoadConfig reads the database configuration from a JSON file
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

// Connect establishes a connection to the database
func Connect(config *DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	return sql.Open("mysql", dsn)
}
