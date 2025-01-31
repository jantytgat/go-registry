package registry

import (
	"database/sql"
	"strings"
)

func Connect(path string) (*sql.DB, error) {
	if !strings.Contains(path, "?") {
		path = path + "?_pragma=foreign_keys(1)"
	} else if !strings.Contains(path, "&_pragma=foreign_keys(1)") {
		path = path + "&_pragma=pragma_foreign_keys(1)"
	}

	return sql.Open("sqlite", path)
}
