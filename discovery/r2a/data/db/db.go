package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type DB interface {
	createTable(tableName string, header []string) error
	insertData(tableName string, header []string, dataChan chan map[string]string) error
}

type SqliteDB struct {
	db *sql.DB
}

func (sdb *SqliteDB) createTable(tableName string, header []string, primaryKey string) error {
	// Create the table
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	for i, column := range header {
		query += fmt.Sprintf("%s TEXT", column)
		if column == primaryKey {
			query += " PRIMARY KEY"
		}
		if i < len(header)-1 {
			query += ", "
		}
	}
	query += ")"
	_, err := sdb.db.Exec(query)
	return err
}

func (sl *SqliteDB) insertData(tableName string, header []string, dataChan chan map[string]string) error {
	// Create a string of column names
	columnNames := strings.Join(header, ", ")

	// Create a string of placeholders
	placeholders := make([]string, len(header))
	for i := range header {
		placeholders[i] = "?"
	}

	// Prepare the insert statement
	stmt, err := sl.db.Prepare(fmt.Sprintf(`
		INSERT INTO %s (%s) VALUES (%s)
	`, tableName, columnNames, strings.Join(placeholders, ", ")))
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each row from the channel
	for row := range dataChan {
		// convert string to interface
		args := make([]interface{}, len(row))
		for i, v := range header {
			args[i] = row[v]
		}

		_, err := stmt.Exec(args...)
		if err != nil {
			return err
		}
	}

	return nil
}
