package registry

import (
	"database/sql"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

type Organization struct {
	Id   int
	Guid string
	Name string
}

func ListOrganizations(r *queryrepo.Repository, db *sql.DB) ([]Organization, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "list"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var orgs []Organization
	for rows.Next() {
		var o Organization
		if err = rows.Scan(&o.Id, &o.Guid, &o.Name); err != nil {
			return nil, err
		}
		orgs = append(orgs, o)
	}
	return orgs, nil
}
