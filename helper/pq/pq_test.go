package postgres_test

import (
	"testing"

	"github.com/khaiql/dbcleaner"
	"github.com/khaiql/dbcleaner/helper/pq"
)

func TestGetTableQuery(t *testing.T) {
	helper := postgres.Helper{}
	query := helper.GetTablesQuery()

	if query != "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public';" {
		t.Error("Wrong query")
	}
}

func TestTruncateTablesCommand(t *testing.T) {
	helper := postgres.Helper{}
	excludedTables := []string{"users"}
	cmd := helper.TruncateTablesCommand(excludedTables)

	if cmd != "TRUNCATE TABLE users" {
		t.Error("Wrong command")
	}
}

func TestInit(t *testing.T) {
	_, err := dbcleaner.FindHelper("postgres")
	if err != nil {
		t.Errorf("Shouldn't get error but got %s", err.Error())
	}
}
