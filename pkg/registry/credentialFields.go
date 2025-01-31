package registry

import (
	"database/sql"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

type CredentialField struct {
	Id           int64
	Name         string
	Value        string
	CredentialId int64
}

func AddCredentialField(r *queryrepo.Repository, db *sql.DB, name, value string, credentialId int64) (CredentialField, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentialFields", "insert"); err != nil {
		return CredentialField{}, err
	}

	var result sql.Result
	if result, err = q.Exec(name, value, credentialId); err != nil {
		return CredentialField{}, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return CredentialField{}, err
	}
	return CredentialField{
		Id:           id,
		Name:         name,
		Value:        value,
		CredentialId: credentialId,
	}, err
}

func DeleteCredentialFieldById(r *queryrepo.Repository, db *sql.DB, id int64) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentialFields", "deleteById"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func GetCredentialFieldById(r *queryrepo.Repository, db *sql.DB, id int64) (CredentialField, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentialFields", "getById"); err != nil {
		return CredentialField{}, err
	}

	var e CredentialField
	return e, q.QueryRow(id).Scan(&e.Id, &e.Name, &e.Value, &e.CredentialId)
}

func ListCredentialFields(r *queryrepo.Repository, db *sql.DB) ([]CredentialField, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentialFields", "list"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentialFields []CredentialField
	for rows.Next() {
		var e CredentialField
		if err = rows.Scan(&e.Id, &e.Name, &e.Value, &e.CredentialId); err != nil {
			return nil, err
		}
		credentialFields = append(credentialFields, e)
	}
	return credentialFields, nil
}

func ListCredentialFieldsByCredentialId(r *queryrepo.Repository, db *sql.DB, id int64) ([]CredentialField, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "credentialFields", "listByCredentialId"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(id); err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentialFields []CredentialField
	for rows.Next() {
		var e CredentialField
		if err = rows.Scan(&e.Id, &e.Name, &e.Value, &e.CredentialId); err != nil {
			return nil, err
		}
		credentialFields = append(credentialFields, e)
	}
	return credentialFields, nil
}
