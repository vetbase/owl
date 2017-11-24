package owl

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	db := Connect("postgres", "postgres://user:pass@localhost/bookstore")
	assert.NotNil(t, db)
}

func TestConnectPanic(t *testing.T) {
	fn := func() {
		Connect("blah", "postgres://user:pass@localhost/bookstore")
	}
	assert.Panics(t, func() { fn() })
}

func TestDatabaseStructIsReturnedOnConnection(t *testing.T) {
	db := Connect("postgres", "postgres://user:pass@localhost/bookstore")
	assert.True(t, reflect.TypeOf(db).String() == "*owl.Database")
}

func TestNewQuery(t *testing.T) {
	q := NewQuery()
	assert.True(t, reflect.TypeOf(q).String() == "*owl.Query")
}
