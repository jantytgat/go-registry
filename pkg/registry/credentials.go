package registry

import (
	"context"
	"database/sql"
)

type Credential struct {
	Id       int64
	Guid     string
	Name     string
	TenantId int64
}

func (r *Registry) DeleteCredentialByGuid(ctx context.Context, guid string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, DeleteByGuid); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteCredentialById(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, DeleteById); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteCredentialByName(ctx context.Context, name string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, DeleteByName); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) GetCredentialByGuid(ctx context.Context, guid string) (Credential, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, GetByGuid); err != nil {
		return Credential{}, err
	}

	var e Credential
	return e, s.QueryRow(guid).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) GetCredentialById(ctx context.Context, id int64) (Credential, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, GetById); err != nil {
		return Credential{}, err
	}

	var e Credential
	return e, s.QueryRow(id).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) GetCredentialByName(ctx context.Context, name string) (Credential, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, GetByName); err != nil {
		return Credential{}, err
	}

	var e Credential
	return e, s.QueryRow(name).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) InsertCredential(ctx context.Context, guid, name string, tenantId int64) (Credential, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, Insert); err != nil {
		return Credential{}, err
	}

	var result sql.Result
	if result, err = s.Exec(guid, name, tenantId); err != nil {
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

func (r *Registry) ListCredentials(ctx context.Context) ([]Credential, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, List); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(); err != nil {
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

func (r *Registry) ListCredentialsByTenantId(ctx context.Context, id int64) ([]Credential, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, CredentialsCollection, ListByTenantId); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(id); err != nil {
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
