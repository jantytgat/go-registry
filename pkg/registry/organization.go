package registry

import (
	"database/sql"
	"fmt"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

type Organization struct {
	Id   int64
	Guid string
	Name string
}

func AddOrganization(r *queryrepo.Repository, db *sql.DB, guid, name string) (Organization, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "insert"); err != nil {
		return Organization{}, err
	}

	var result sql.Result
	if result, err = q.Exec(guid, name); err != nil {
		return Organization{}, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return Organization{}, err
	}
	return Organization{
		Id:   id,
		Guid: guid,
		Name: name,
	}, err
}

func DeleteOrganizationByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "deleteByGuid"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteOrganizationById(r *queryrepo.Repository, db *sql.DB, id int64) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "deleteById"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteOrganizationByName(r *queryrepo.Repository, db *sql.DB, name string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "deleteByName"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func GetOrganizationByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (Organization, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "getByGuid"); err != nil {
		return Organization{}, err
	}

	var o Organization
	return o, q.QueryRow(guid).Scan(&o.Id, &o.Guid, &o.Name)
}

func GetOrganizationById(r *queryrepo.Repository, db *sql.DB, id int64) (Organization, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "getById"); err != nil {
		return Organization{}, err
	}

	var o Organization
	return o, q.QueryRow(id).Scan(&o.Id, &o.Guid, &o.Name)
}

func GetOrganizationByName(r *queryrepo.Repository, db *sql.DB, name string) (Organization, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "getByName"); err != nil {
		return Organization{}, err
	}

	var o Organization
	return o, q.QueryRow(name).Scan(&o.Id, &o.Guid, &o.Name)
}

func ListOrganizations(r *queryrepo.Repository, db *sql.DB) ([]Organization, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "organizations", "list"); err != nil {
		fmt.Println("Error preparing listing organizations")
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(); err != nil {
		fmt.Println("Error querying listing organizations")
		return nil, err
	}
	defer rows.Close()

	var organizations []Organization
	for rows.Next() {
		var o Organization
		if err = rows.Scan(&o.Id, &o.Guid, &o.Name); err != nil {
			fmt.Println("Error scanning listing organizations")
			return nil, err
		}
		organizations = append(organizations, o)
	}
	return organizations, nil
}
