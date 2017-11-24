package owl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableSetsCorrectName(t *testing.T) {
	db := Connect("postgres", "postgres://user:pass@localhost/bookstore")

	db.Table("users")

	assert.Equal(t, "users", db.Query.Table)
}
