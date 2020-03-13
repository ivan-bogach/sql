package sql

import (
	"database/sql"
	"strings"

	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbType = "sqlite3"
)

func prepareExec(db *sql.DB, queryString string) error {
	statement, err := db.Prepare(queryString)
	if err != nil {
		color.Red("prepareExec():")
		return err
	}
	statement.Exec()
	return nil
}

func getColumns(columnsWithType []string) string {
	var result string
	for _, columnWithType := range columnsWithType {
		result += columnWithType + ","
	}
	return strings.TrimSuffix(result, ",")
}

// CreateTableIfNotExists creates table with columns
func CreateTableIfNotExists(columnsWithType []string, tableName, dbPath string) error {
	queryString := "CREATE TABLE IF NOT EXISTS " + tableName + " (" + getColumns(columnsWithType) + ");"
	db, _ := sql.Open(dbType, dbPath)
	defer db.Close()
	err := prepareExec(db, queryString)
	if err != nil {
		return err
	}
	return err
}
