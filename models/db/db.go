package db

import (
	"database/sql"
	"log"

	// register sqlite
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	initDB()
}

// Row is wrapper for sql.Row
type Row struct {
	*sql.Row
}

// Rows is wrapper for sql.Rows
type Rows struct {
	*sql.Rows
}

// Scan scan value from row
func (r *Row) Scan(dest ...interface{}) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Panicf("error while scanning row: %v\n", err)
	}
	return nil
}

// Scan scan value from rows
func (rs *Rows) Scan(dest ...interface{}) error {
	err := rs.Rows.Scan(dest...)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Panicf("error while scanning rows: %v\n", err)
	}
	return nil
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "appleforum.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// QueryRow query one row
func QueryRow(query string, args ...interface{}) *Row {
	return &Row{db.QueryRow(query, args...)}
}

// Query query multiple rows
func Query(query string, args ...interface{}) *Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Panicf("error with sql query: %v, msg: %v", query, err)
	}
	return &Rows{rows}
}

// Exec executes a query
func Exec(query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Panicf("err with sql exec: %v, msg: %v", query, err)
	}
}
