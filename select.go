package owl

import "strings"

// Select builds a select statement
func (db *Database) Select(table string, cols []string) *Database {
	db.Query = NewQuery()
	db.Query.Action = "SELECT"
	db.Query.Columns = columns(cols)
	db.Query.Table = table

	return db
}

// Where sets the where condition for the sql query
func (db *Database) Where(condition string, args ...string) *Database {
	db.Query.Where = "WHERE"
	db.Query.WhereCondition = condition
	db.Query.Arguments = args

	return db
}

func columns(cols []string) string {
	var c string
	if len(cols) > 0 {
		c = strings.Join(cols, ",")
	} else {
		c = "*"
	}

	return c
}
