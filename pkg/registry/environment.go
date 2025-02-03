package registry

import (
	"context"
	"database/sql"
)

type Environment struct {
	Id       int64
	Guid     string
	Name     string
	TenantId int64
}

func (r *Registry) DeleteEnvironmentByGuid(ctx context.Context, guid string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, DeleteByGuid); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteEnvironmentById(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, DeleteById); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteEnvironmentByNameAndTenantId(ctx context.Context, name string, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, DeleteByNameAndTenantId); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(name, id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteEnvironmentByNameAndTenantName(ctx context.Context, name, tenantName, organizationName string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, DeleteByNameAndTenantNameAndOrganizationName); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(name, tenantName, organizationName); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) GetEnvironmentByGuid(ctx context.Context, guid string) (Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, GetByGuid); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, s.QueryRow(guid).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) GetEnvironmentById(ctx context.Context, id int64) (Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, GetById); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, s.QueryRow(id).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) GetEnvironmentByNameAndTenantId(ctx context.Context, name string, id int64) (Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, GetByNameAndTenantId); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, s.QueryRow(name, id).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) GetEnvironmentByNameAndTenantNameAndOrganizationName(ctx context.Context, name, tenant, organization string) (Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, GetByNameAndTenantNameAndOrganizationName); err != nil {
		return Environment{}, err
	}

	var e Environment
	return e, s.QueryRow(name, tenant, organization).Scan(&e.Id, &e.Guid, &e.Name, &e.TenantId)
}

func (r *Registry) InsertEnvironment(ctx context.Context, guid, name string, tenantId int64) (Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, Insert); err != nil {
		return Environment{}, err
	}

	var result sql.Result
	if result, err = s.Exec(guid, name, tenantId); err != nil {
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

func (r *Registry) ListEnvironments(ctx context.Context) ([]Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, List); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(); err != nil {
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

func (r *Registry) ListEnvironmentsByTenantId(ctx context.Context, id int64) ([]Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, ListByTenantId); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(id); err != nil {
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

func (r *Registry) ListEnvironmentsByTenantNameAndOrganizationName(ctx context.Context, tenant, organization string) ([]Environment, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, EnvironmentsCollection, ListByTenantNameAndOrganizationName); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(tenant, organization); err != nil {
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
