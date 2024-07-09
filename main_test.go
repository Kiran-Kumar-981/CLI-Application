//kindly first give the correct details about the database
package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	testDSN = "UserName:Password@tcp(127.0.0.1:3306)/mydata" // Replace with your actual test database DSN
)

// TestDBConnection verifies the database connection and basic functionality
func TestDBConnection(t *testing.T) {
	// Attempt to open a connection to the database
	db, err := sql.Open("mysql", testDSN)
	if err != nil {
		t.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	// Check if the database connection is alive
	err = db.Ping()
	if err != nil {
		t.Fatalf("Error pinging database: %v", err)
	}

	t.Logf("Successfully connected to the database with DSN: %s", testDSN)

	// Perform additional tests if needed, like querying or inserting test data
}
