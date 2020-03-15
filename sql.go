package sql

import (
	"database/sql"
	"log"
	"strings"

	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbType = "sqlite3"
)

func prepareExec(db *sql.DB, queryString string) {
	c := color.New(color.FgGreen)
	c.Printf("SQL Query: ' %s '\n ", queryString)

	statement, err := db.Prepare(queryString)
	if err != nil {
		color.Red("Error occurred")
		log.Fatal(err)
	}
	statement.Exec()
	d := color.New(color.FgGreen, color.Bold)

	d.Println("Ok!")

}

func CreateTableIfNotExists(tableName, tablePath, columnsString string) {
	queryString := "CREATE TABLE IF NOT EXISTS " + tableName + " (" + columnsString + ");"
	db, _ := sql.Open(dbType, tablePath)
	defer db.Close()
	prepareExec(db, queryString)

}

func SelectRow(tablePath, queryString string) *sql.Row {
	c := color.New(color.FgGreen)
	c.Printf("SQL Query: ' %s '\n ", queryString)
	db, _ := sql.Open("sqlite3", "db/db.sqlite")
	defer db.Close()

	row := db.QueryRow(queryString)
	d := color.New(color.FgGreen, color.Bold)
	d.Println("Ok!")

	return row
}

func SelectRows(tablePath, queryString string) *sql.Rows {
	c := color.New(color.FgGreen)
	c.Printf("SQL Query: ' %s '\n ", queryString)

	db, _ := sql.Open(dbType, tablePath)
	defer db.Close()
	rows, err := db.Query(queryString)
	if err != nil {
		color.Red("Error occurred")
	}
	d := color.New(color.FgGreen, color.Bold)
	d.Println("Ok!")
	return rows
}

func ExecRow(tablePath, queryString string) {
	db, _ := sql.Open(dbType, tablePath)
	defer db.Close()
	prepareExec(db, queryString)
}

func ExecRows(tablePath, queryString string) {
	c := color.New(color.FgGreen)
	c.Printf("SQL Query: ' %s '\n ", queryString)

	db, _ := sql.Open(dbType, tablePath)
	defer db.Close()

	result, err := db.Exec(queryString)
	if err != nil {
		color.Red("Error occurred")
		log.Fatal(err)
	}
	numRows, err := result.RowsAffected()
	if err != nil {
		color.Red("Error occurred")
		log.Fatal(err)
	}

	d := color.New(color.FgGreen, color.Bold)
	d.Println("Added %d rows!", numRows)
}

func CommaJoin(sl []string) string {
	var result string
	for _, val := range sl {
		result += val + ","
	}
	result = strings.TrimSuffix(result, ",")
	return result

}
