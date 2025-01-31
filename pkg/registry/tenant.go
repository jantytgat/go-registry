package registry

import (
	"database/sql"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

type Tenant struct {
	Id             int64
	Guid           string
	Name           string
	OrganizationId int64
}

func AddTenant(r *queryrepo.Repository, db *sql.DB, guid, name string, organizationId int64) (Tenant, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "insert"); err != nil {
		return Tenant{}, err
	}

	var result sql.Result
	if result, err = q.Exec(guid, name, organizationId); err != nil {
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

func DeleteTenantByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "deleteByGuid"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(guid); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteTenantById(r *queryrepo.Repository, db *sql.DB, id int64) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "deleteById"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteTenantByName(r *queryrepo.Repository, db *sql.DB, name string) (int64, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "deleteByName"); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = q.Exec(name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func GetTenantByGuid(r *queryrepo.Repository, db *sql.DB, guid string) (Tenant, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "getByGuid"); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, q.QueryRow(guid).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func GetTenantById(r *queryrepo.Repository, db *sql.DB, id int64) (Tenant, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "getById"); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, q.QueryRow(id).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func GetTenantByName(r *queryrepo.Repository, db *sql.DB, name string) (Tenant, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "getByName"); err != nil {
		return Tenant{}, err
	}

	var t Tenant
	return t, q.QueryRow(name).Scan(&t.Id, &t.Guid, &t.Name, &t.OrganizationId)
}

func ListTenants(r *queryrepo.Repository, db *sql.DB) ([]Tenant, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "list"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(); err != nil {
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

func ListTenantsByOrganizationId(r *queryrepo.Repository, db *sql.DB, id int64) ([]Tenant, error) {
	var err error

	var q *sql.Stmt
	if q, err = r.DbPrepare(db, "tenants", "listByOrganizationId"); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = q.Query(id); err != nil {
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
