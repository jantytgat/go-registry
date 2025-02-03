package registry

import (
	"context"
	"database/sql"
)

type CredentialField struct {
	Id           int64
	Name         string
	Value        string
	CredentialId int64
}

func (r *Registry) DeleteCredentialFieldByCredentialId(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialFieldsCollection, DeleteByCredentialId); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteCredentialFieldById(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialFieldsCollection, DeleteById); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) GetCredentialFieldById(ctx context.Context, id int64) (CredentialField, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialFieldsCollection, GetById); err != nil {
		return CredentialField{}, err
	}

	var e CredentialField
	return e, s.QueryRow(id).Scan(&e.Id, &e.Name, &e.Value, &e.CredentialId)
}

func (r *Registry) InsertCredentialField(ctx context.Context, name, value string, credentialId int64) (CredentialField, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialFieldsCollection, Insert); err != nil {
		return CredentialField{}, err
	}

	var result sql.Result
	if result, err = s.Exec(name, value, credentialId); err != nil {
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

func (r *Registry) ListCredentialFields(ctx context.Context) ([]CredentialField, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialFieldsCollection, List); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(); err != nil {
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

func (r *Registry) ListCredentialFieldsByCredentialId(ctx context.Context, id int64) ([]CredentialField, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialFieldsCollection, ListByCredentialId); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(id); err != nil {
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
