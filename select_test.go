package owl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectNoColumns(t *testing.T) {
	db := Connect("postgres", "postgres://user:pass@localhost/bookstore")
	defer db.Close()

	db.Table("users").Select()
	assert.Equal(t, "SELECT*FROMusers", db.Query.SQL())
}

func TestSelectWithColumns(t *testing.T) {
	db := Connect("postgres", "postgres://user:pass@localhost/bookstore")

	db.Table("users").Select("name", "email")
	assert.Equal(t, "SELECT name,email FROMusers", db.Query.SQL())
}
