package db

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Table    string `json:"table"`
}

const rootCert = `
-----BEGIN CERTIFICATE-----
MIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j
b20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB
CSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97
nh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt
43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P
T19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4
gdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO
BgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR
TLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw
DQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr
hMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg
06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF
PnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls
YSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk
CAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=
-----END CERTIFICATE-----
`

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
	// new certificate pool and root certificate
	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM([]byte(rootCert)); !ok {
		return nil, fmt.Errorf("failed to append root certificate")
	}

	// TLS configuration
	tlsConfig := &tls.Config{
		RootCAs: rootCertPool,
	}

	// Register the TLS configuration
	err := mysql.RegisterTLSConfig("azuremysql", tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to register TLS config: %v", err)
	}

	// Construct the DSN using TLS
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=azuremysql",
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

	fmt.Println("Successfully connected to the database using TLS")
	return db, nil
}
