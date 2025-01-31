package registry

import (
	"database/sql"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

type Credential struct {
	Id       int64
	Guid     string
	Name     string
	TenantId int64
}

func AddCredential(r *queryrepo.Repository, db *sql.DB, guid, name string, tenantId int64) (Credential, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "insert"); err != nil {
		return Credential{}, err
	}

	var result sql.Result
	if result, err = q.Exec(guid, name, tenantId); err != nil {
		return Credential{}, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return Credential{}, err
	}
	return Credential{
		Id:       id,
		Guid:     guid,
		Name:     name,
		TenantId: tenantId,
	}, err
}

func DeleteCredentialByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "deleteByGuid"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteCredentialById(r *queryrepo.Repository, db *sql.DB, id int64) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "deleteById"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteCredentialByName(r *queryrepo.Repository, db *sql.DB, name string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "deleteByName"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func GetCredentialByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (Credential, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "getByGuid"); err != nil {
		return Credential{}, err
	}

	var e Credential
	return e, q.QueryRow(guid).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func GetCredentialById(r *queryrepo.Repository, db *sql.DB, id int64) (Credential, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "getById"); err != nil {
		return Credential{}, err
	}

	var e Credential
	return e, q.QueryRow(id).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func GetCredentialByName(r *queryrepo.Repository, db *sql.DB, name string) (Credential, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "getByName"); err != nil {
		return Credential{}, err
	}

	var e Credential
	return e, q.QueryRow(name).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func ListCredentials(r *queryrepo.Repository, db *sql.DB) ([]Credential, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "list"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []Credential
	for rows.Next() {
		var e Credential
		if err = rows.Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId); err != nil {
			return nil, err
		}
		credentials = append(credentials, e)
	}
	return credentials, nil
}

func ListCredentialsByTenantId(r *queryrepo.Repository, db *sql.DB, id int64) ([]Credential, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentials", "listByTenantId"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(id); err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []Credential
	for rows.Next() {
		var e Credential
		if err = rows.Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId); err != nil {
			return nil, err
		}
		credentials = append(credentials, e)
	}
	return credentials, nil
}
