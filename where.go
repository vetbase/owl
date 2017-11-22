package owl

// Where sets the where condition for the sql query
func (db *Database) Where(condition string, args ...string) *Database {
	db.Query.Where = "WHERE"
	db.Query.WhereCondition = condition
	db.Query.Arguments = args

	return db
}

// WhereOr sets a or where condition for the sql query
func (db *Database) WhereOr(condition string, args ...string) *Database {
	db.Query.Where = "OR WHERE"
	db.Query.WhereCondition = condition
	db.Query.Arguments = args

	return db
}
