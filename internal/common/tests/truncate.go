package tests

import (
	"database/sql"
	"fmt"
	"testing"
)

func TruncateTables(t *testing.T, db *sql.DB, tableNames ...string) {
	t.Helper()
	for _, tbn := range tableNames {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE %s", tbn))
		if err != nil {
			t.Fatalf("Fail to truncate table %s: %v", tbn, err)
		}
	}
}
