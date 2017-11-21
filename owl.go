package owl

import (
	"bytes"
	"database/sql"
	// "database/sql/driver"
)

// Database contains the connection details
type Database struct {
	Driver     string
	Connection *sql.DB
	Query      *Query
}

// Query holds information about SQL being executed
type Query struct {
	Action         string
	Columns        string
	Table          string
	Where          string
	WhereCondition string
	Arguments      []string
}

// NewDatabase creates a new database struct
func NewDatabase(driver string, db *sql.DB) *Database {
	return &Database{
		Driver:     driver,
		Connection: db,
	}
}

// NewQuery creates a new query to be executed
func NewQuery() *Query {
	return &Query{}
}

// Connect will attempt to connect to specified database
func Connect(driver string, dns string) *Database {
	db, err := sql.Open(driver, dns)
	if err != nil {
		panic(err)
	}
	return NewDatabase(driver, db)
}

// Execute query on database
func (db *Database) Execute() {
	sql := db.Query.SQL()
	db.Connection.Exec(sql, db.Query.Arguments)
}

// SQL will return a string representation of the current query
func (q *Query) SQL() string {
	var buffer bytes.Buffer
	buffer.WriteString(q.Action)
	buffer.WriteString(q.Columns)
	buffer.WriteString("FROM")
	buffer.WriteString(q.Table)
	buffer.WriteString(q.Where)
	buffer.WriteString(q.WhereCondition)
	return buffer.String()
}
