package interfaces

import (
	"database/sql"
	"time"
)

// ISQL is an abstract for sql database
type ISQL interface {
	OpenConnection()
	GetDBMaster() (*sql.DB, error)
	GetDBSlave() (*sql.DB, error)
	SetConnMaxLifetime(time.Duration)
	SetMaxIdleConns(int)
	SetMaxOpenConns(int)
}
