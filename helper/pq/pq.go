package postgres

import (
	"fmt"
	"strings"

	"github.com/khaiql/dbcleaner"
)

type Helper struct{}

func init() {
	dbcleaner.RegisterHelper("postgres", Helper{})
}

const queryAllTables = "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public';"

func (Helper) GetTablesQuery() string {
	return queryAllTables
}

func (Helper) TruncateTablesCommand(tableNames []string) string {
	if len(tableNames) == 0 {
		return ""
	} else {
		return fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tableNames, ","))
	}
}
