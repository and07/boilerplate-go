package data

import (
	"fmt"
	"time"

	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewConnection creates the connection to the database
func NewConnection(config *utils.Configurations, logger hclog.Logger) (*sqlx.DB, error) {

	var conn string

	if config.DBConn != "" {
		conn = config.DBConn
	} else {
		host := config.DBHost
		port := config.DBPort
		user := config.DBUser
		dbName := config.DBName
		password := config.DBPass
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	}
	logger.Debug("connection string", conn)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	//TODO
	db.SetMaxOpenConns(25) // Too many connections
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute)

	return db, nil
}
