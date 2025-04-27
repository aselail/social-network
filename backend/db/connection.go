package db

import (
	"database/sql"
	"fmt"
)

// Database struct manages the database connection and table name.
type Database struct {
	db *sql.DB
}

var Connection Database

// Open a new database connection
func (d *Database) Open(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	d.db = db
	return nil
}

// Close closes the database connection.
func (d *Database) Close() error {
	return d.db.Close()
}
