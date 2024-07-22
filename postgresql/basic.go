// Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package flamingos

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

// Reading PostgreSQL database connection parameters
//
//	file:    The name of the TOML configuration file containing the path
//	returns: Configuration content, error messages
func GetPsqlConfig(file string) (string, string) {
	config, err := toml.LoadFile(file)
	if err != nil {
		return "", err.Error()
	}
	var str = ""
	var paramKey = [6]string{"host", "port", "user", "password", "dbname",
		"sslmode"}
	// The parameters used to connect to a PostgreSQL database are a string of
	// space-separated key-value pairs, so you don't need to read the parameters
	// in the TOML file one by one, you can just splice them together into a
	// string.
	for i := 0; i < 6; i++ {
		// Check if the parameter exists.
		err := config.Get(paramKey[i])
		if err == nil {
			return "", "Parameters are missing from the configuration file:" +
				paramKey[i]
		}
		str = str + paramKey[i] + "=" + config.Get(paramKey[i]).(string) + " "
	}
	return str, ""
}

// Connecting to a PostgreSQL Database
//
//	config: database connection parameters
//	returns: Database connection handle, error message
func ConnectPsql(config string) (*sql.DB, string) {
	// The sql.Open() function obviously does only a limited amount of testing
	// when determining whether the connection was successful, i.e., it only
	// tests whether the first argument is correct. In the case that the first
	// parameter is correct, no error will be reported regardless of the content
	// of the second string parameter (even if it is null). So further testing
	// needs to be done to accurately determine if the connection was successful.
	db, err := sql.Open("postgres", config)
	if err != nil {
		return nil, "An error occurred while connecting to the database\n" +
			err.Error()
	}
	_, err = db.Query("SELECT 0")
	if err != nil {
		return nil, "An error occurred while querying the database\n" + err.Error()
	}
	return db, ""
}

// Close the database connection
//
//	db: database connection handle
//	returns: error message
func ClosePsql(db *sql.DB) string {
	err := db.Close()
	if err != nil {
		return "Error while closing database connection\n" + err.Error()
	}
	return ""
}

// Query Data
//
//	db: database connection handle
//	sel: Query statement. The statement can end with or without a semicolon.
//	returns: query result, error value
func PsqlSelect(db *sql.DB, sel string) (*sql.Rows, string) {
	rows, err := db.Query(sel)
	if err != nil {
		return nil, err.Error()
	}
	return rows, ""
}

// Insertion, modification and deletion of data
//
//	db: database connection handle
//	exec: sql statement. The statement can end with or without a semicolon.
//	returns: error value
func PsqlExec(db *sql.DB, exec string) string {
	_, err := db.Exec(exec)
	if err != nil {
		return "An error occurred while executing '" + exec + "'\n" + err.Error()
	}
	return ""
}
