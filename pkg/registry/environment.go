package registry

import (
	"database/sql"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

type Environment struct {
	Id       int64
	Guid     string
	Name     string
	TenantId int64
}

func AddEnvironment(r *queryrepo.Repository, db *sql.DB, guid, name string, tenantId int64) (Environment, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "insert"); err != nil {
		return Environment{}, err
	}

	var result sql.Result
	if result, err = q.Exec(guid, name, tenantId); err != nil {
		return Environment{}, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return Environment{}, err
	}
	return Environment{
		Id:       id,
		Guid:     guid,
		Name:     name,
		TenantId: tenantId,
	}, err
}

func DeleteEnvironmentByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "deleteByGuid"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteEnvironmentById(r *queryrepo.Repository, db *sql.DB, id int64) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "deleteById"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteEnvironmentByName(r *queryrepo.Repository, db *sql.DB, name string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "deleteByName"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func GetEnvironmentByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (Environment, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "getByGuid"); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, q.QueryRow(guid).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func GetEnvironmentById(r *queryrepo.Repository, db *sql.DB, id int64) (Environment, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "getById"); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, q.QueryRow(id).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func GetEnvironmentByName(r *queryrepo.Repository, db *sql.DB, name string) (Environment, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "getByName"); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, q.QueryRow(name).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func ListEnvironments(r *queryrepo.Repository, db *sql.DB) ([]Environment, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "list"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var environments []Environment
	for rows.Next() {
		var e Environment
		if err = rows.Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId); err != nil {
			return nil, err
		}
		environments = append(environments, e)
	}
	return environments, nil
}

func ListEnvironmentsByTenantId(r *queryrepo.Repository, db *sql.DB, id int64) ([]Environment, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "environments", "listByTenantId"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(id); err != nil {
		return nil, err
	}
	defer rows.Close()

	var environments []Environment
	for rows.Next() {
		var e Environment
		if err = rows.Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId); err != nil {
			return nil, err
		}
		environments = append(environments, e)
	}
	return environments, nil
}
