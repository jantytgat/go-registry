package registry

import (
	"context"
	"database/sql"
	"fmt"
)

type Organization struct {
	Id   int64
	Guid string
	Name string
}

func (r *Registry) DeleteOrganizationByGuid(ctx context.Context, guid string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, DeleteByGuid); err != nil {
		return 0, err
	}

	var res sql.Result
	if res, err = s.ExecContext(ctx, guid); err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *Registry) DeleteOrganizationById(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, DeleteById); err != nil {
		return 0, err
	}

	var res sql.Result
	if res, err = s.ExecContext(ctx, id); err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *Registry) DeleteOrganizationByName(ctx context.Context, name string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, DeleteByName); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) GetOrganizationByGuid(ctx context.Context, guid string) (Organization, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, GetByGuid); err != nil {
		return Organization{}, err
	}

	var o Organization
	return o, s.QueryRowContext(ctx, guid).Scan(&o.Id, &o.Guid, &o.Name)
}

func (r *Registry) GetOrganizationById(ctx context.Context, id int64) (Organization, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, GetById); err != nil {
		return Organization{}, err
	}

	var o Organization
	return o, s.QueryRowContext(ctx, id).Scan(&o.Id, &o.Guid, &o.Name)
}

func (r *Registry) GetOrganizationByName(ctx context.Context, name string) (Organization, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, GetByName); err != nil {
		return Organization{}, err
	}

	var o Organization
	return o, s.QueryRowContext(ctx, name).Scan(&o.Id, &o.Guid, &o.Name)
}

func (r *Registry) InsertOrganization(ctx context.Context, guid, name string) (Organization, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, Insert); err != nil {
		return Organization{}, err
	}

	var result sql.Result
	if result, err = s.ExecContext(ctx, guid, name); err != nil {
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

func (r *Registry) ListOrganizations(ctx context.Context) ([]Organization, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, OrganizationsCollection, List); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.QueryContext(ctx); err != nil {
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
