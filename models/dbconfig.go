package models

type DbConfig struct {
	Host       string
	User       string
	Password   string
	DBName     string
	Port       string
	SSLEnabled string `json:"ssl_mode"`
}

// type sslMode string

// const (
// 	SSLEnabled  sslMode = "enabled"
// 	SSLDisabled sslMode = "disabled"
// )

func NewConfig(host, user, pass, dbName, port, sslMode string) *DbConfig {
	return &DbConfig{host, user, pass, dbName, port, sslMode}
}
