package owl

// Table creates new query and sets table name
func (db *Database) Table(name string) *Database {
	db.Query = NewQuery()
	db.Query.Table = name

	return db
}
