package adapter

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const connStringTemplate = "%s:%s@tcp(%s)/%s?timeout=%v&charset=%s&parseTime=true&loc=Local"

type (
	MySQLConfig struct {
		Host            string
		User            string
		Password        string
		Name            string
		Timeout         time.Duration
		Charset         string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime time.Duration
	}

	MySQLConnector struct {
		config *MySQLConfig
		query  map[string]string
		db     *sql.DB
		dbMu   sync.Mutex
	}
)

// MySQLAdapter mysql contract
type MySQLAdapter interface {
	Query(ctx context.Context, query string, args ...interface{}) *sql.Row
	Queries(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Ping(ctx context.Context) error
}

// NewMySQLConnector is the create new mysql connection
func NewMySQLConnector(cnf *MySQLConfig) *MySQLConnector {
	conn := &MySQLConnector{
		config: cnf,
	}

	conn.Open()

	return conn
}

// Open session mysql database
func (i *MySQLConnector) Open() *sql.DB {
	connectionString := fmt.Sprintf(connStringTemplate, i.config.User, i.config.Password, i.config.Host, i.config.Name, i.config.Timeout, i.config.Charset)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(i.config.ConnMaxLifetime)
	db.SetMaxIdleConns(i.config.MaxIdleConns)
	db.SetMaxOpenConns(i.config.MaxOpenConns)

	i.addDB(db)

	return db

}

// addDB create database
func (i *MySQLConnector) addDB(db *sql.DB) {
	i.dbMu.Lock()
	defer i.dbMu.Unlock()
	i.db = db
}

// check connection database
func (i *MySQLConnector) check() error {

	if i.db == nil {
		i.Open()
	}

	return i.db.Ping()
}

// Query executes a query that returns row, typically a SELECT.
func (i *MySQLConnector) Query(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {

	i.check()
	return i.db.QueryRowContext(ctx, query, args...)

}

// Queries executes a query that returns rows, typically a SELECT.
func (i *MySQLConnector) Queries(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {

	if err = i.check(); err != nil {
		err = err
		return
	}

	rows, err = i.db.QueryContext(ctx, query, args...)

	return
}

// BeginTx start a transaction database
func (i *MySQLConnector) BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sql.Tx, err error) {
	tx, err = i.db.BeginTx(ctx, opts)
	return
}

// Exec executes a query without returning any rows.
func (i *MySQLConnector) Exec(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	if err = i.check(); err != nil {
		return
	}
	result, err = i.db.ExecContext(ctx, query, args...)

	return
}

// Ping database connection
func (i *MySQLConnector) Ping(ctx context.Context) error {
	return i.db.PingContext(ctx)
}
