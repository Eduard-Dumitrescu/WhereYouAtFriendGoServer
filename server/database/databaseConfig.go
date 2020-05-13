package database

import (
	"fmt"
)

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig sets configuration data
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "admin",
		DBName:   "WhereYouAtFriend",
	}
	return &dbConfig
}

// DbURL generates db connection url
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

// DbURLDefault return db url with default config
func DbURLDefault() string {
	return DbURL(BuildDBConfig())
}
