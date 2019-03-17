package infrastructures

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//SQLInfrastructure DBslave or master
type SQLInfrastructure struct {
	SQLSlave  SQL
	SQLMaster SQL
}

// SQL connection attributes
type SQL struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Charset  string
	DB       *sql.DB
}

var urlFormat = "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local"

// OpenConnection gets a handle for a database
func (s *SQLInfrastructure) OpenConnection() {
	// Get database handler master
	dataSourceMaster := fmt.Sprintf(urlFormat, s.SQLMaster.User, s.SQLMaster.Password,
		s.SQLMaster.Host, s.SQLMaster.Port, s.SQLMaster.DBName, s.SQLMaster.Charset)
	dbMaster, err := sql.Open("mysql", dataSourceMaster)
	if err != nil {
		panic(fmt.Errorf("failed dial database master: %s", err.Error()))
	}
	s.SQLMaster.DB = dbMaster

	// Get database handler slave
	dataSourceSlave := fmt.Sprintf(urlFormat, s.SQLSlave.User, s.SQLSlave.Password,
		s.SQLSlave.Host, s.SQLSlave.Port, s.SQLSlave.DBName, s.SQLSlave.Charset)
	dbSlave, err := sql.Open("mysql", dataSourceSlave)
	if err != nil {
		panic(fmt.Errorf("failed dial database slave : %s", err.Error()))
	}
	s.SQLSlave.DB = dbSlave
}

// GetDBMaster gets master database connection
func (s *SQLInfrastructure) GetDBMaster() (*sql.DB, error) {
	err := s.SQLMaster.DB.Ping()
	if err != nil {
		return nil, err
	}

	return s.SQLMaster.DB, nil
}

// GetDBSlave gets slave database connection
func (s *SQLInfrastructure) GetDBSlave() (*sql.DB, error) {
	err := s.SQLSlave.DB.Ping()
	if err != nil {
		return nil, err
	}

	return s.SQLSlave.DB, nil
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (s *SQLInfrastructure) SetConnMaxLifetime(connMaxLifetime time.Duration) {
	s.SQLMaster.DB.SetConnMaxLifetime(connMaxLifetime)
	s.SQLSlave.DB.SetConnMaxLifetime(connMaxLifetime)
}

// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool.
func (s *SQLInfrastructure) SetMaxIdleConns(maxIdleConn int) {
	s.SQLMaster.DB.SetMaxIdleConns(maxIdleConn)
	s.SQLSlave.DB.SetMaxIdleConns(maxIdleConn)
}

// SetMaxOpenConns sets the maximum amount of time a connection may be reused.
func (s *SQLInfrastructure) SetMaxOpenConns(maxOpenConn int) {
	s.SQLMaster.DB.SetMaxOpenConns(maxOpenConn)
	s.SQLSlave.DB.SetMaxOpenConns(maxOpenConn)
}
