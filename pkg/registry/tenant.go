package registry

import (
	"context"
	"database/sql"
)

type Tenant struct {
	Id             int64
	Guid           string
	Name           string
	OrganizationId int64
}

func (r *Registry) DeleteTenantByGuid(ctx context.Context, guid string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, DeleteByGuid); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteTenantById(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, DeleteById); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteTenantByNameAndOrganizationId(ctx context.Context, name string, organizationId int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, DeleteByNameAndOrganizationId); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, name, organizationId); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteTenantByNameAndOrganizationName(ctx context.Context, name, organization string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, DeleteByNameAndOrganizationName); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, name, organization); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) GetTenantByGuid(ctx context.Context, guid string) (Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, GetByGuid); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, s.QueryRowContext(ctx, guid).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func (r *Registry) GetTenantById(ctx context.Context, id int64) (Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, GetById); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, s.QueryRowContext(ctx, id).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func (r *Registry) GetTenantByNameAndOrganizationId(ctx context.Context, name string, id int64) (Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, GetByNameAndOrganizationId); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, s.QueryRowContext(ctx, name, id).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func (r *Registry) GetTenantByNameAndOrganizationName(ctx context.Context, name, organization string) (Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, GetByNameAndOrganizationName); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, s.QueryRowContext(ctx, name, organization).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func (r *Registry) InsertTenant(ctx context.Context, guid, name string, organizationId int64) (Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, Insert); err != nil {
		return Tenant{}, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, guid, name, organizationId); err != nil {
		return Tenant{}, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return Tenant{}, err
	}
	return Tenant{
		Id:             id,
		Guid:           guid,
		Name:           name,
		OrganizationId: organizationId,
	}, err
}

func (r *Registry) ListTenants(ctx context.Context) ([]Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, List); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.QueryContext(ctx); err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []Tenant
	for rows.Next() {
		var t Tenant
		if err = rows.Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId); err != nil {
			return nil, err
		}
		tenants = append(tenants, t)
	}
	return tenants, nil
}

func (r *Registry) ListTenantsByOrganizationId(ctx context.Context, id int64) ([]Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, ListByOrganizationId); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.QueryContext(ctx, id); err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []Tenant
	for rows.Next() {
		var t Tenant
		if err = rows.Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId); err != nil {
			return nil, err
		}
		tenants = append(tenants, t)
	}
	return tenants, nil
}

func (r *Registry) ListTenantsByOrganizationName(ctx context.Context, organization string) ([]Tenant, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, TenantsCollection, ListByOrganizationName); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.QueryContext(ctx, organization); err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []Tenant
	for rows.Next() {
		var t Tenant
		if err = rows.Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId); err != nil {
			return nil, err
		}
		tenants = append(tenants, t)
	}
	return tenants, nil
}
