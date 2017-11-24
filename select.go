package owl

import "strings"

// Select builds a select statement
func (db *Database) Select(cols ...string) *Database {
	db.Query.Action = "SELECT"
	db.Query.Columns = columns(cols)

	return db
}

func columns(cols []string) string {
	c := "*"
	if len(cols) > 0 {
		c = " " + strings.Join(cols, ",") + " "
	}

	return c
}
