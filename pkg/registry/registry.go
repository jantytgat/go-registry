package registry

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"sync"

	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
)

const (
	RawCollection = "raw"
)

// Collections
const (
	ConnectionsCollection      = "connections"
	ConnectionTypesCollection  = "connection_types"
	CredentialsCollection      = "credentials"
	CredentialFieldsCollection = "credential_fields"
	EnvironmentsCollection     = "environments"
	OrganizationsCollection    = "organizations"
	TenantsCollection          = "tenants"
)

// Queries
const (
	DeleteByCredentialId                         = "deleteByCredentialId"
	DeleteByGuid                                 = "deleteByGuid"
	DeleteById                                   = "deleteById"
	DeleteByName                                 = "deleteByName"
	DeleteByNameAndOrganizationId                = "deleteByNameAndOrganizationId"
	DeleteByNameAndOrganizationName              = "deleteByNameAndOrganizationName"
	DeleteByNameAndTenantId                      = "deleteByNameAndTenantId"
	DeleteByNameAndTenantNameAndOrganizationName = "deleteByNameAndTenantNameAndOrganizationName"
	GetByGuid                                    = "getByGuid"
	GetById                                      = "getById"
	GetByName                                    = "getByName"
	GetByNameAndTenantId                         = "getByNameAndTenantId"
	GetByNameAndTenantNameAndOrganizationName    = "getByNameAndTenantNameAndOrganizationName"
	GetByNameAndOrganizationId                   = "getByNameAndOrganizationId"
	GetByNameAndOrganizationName                 = "getByNameAndOrganizationName"
	Insert                                       = "insert"
	InsertWithOrganizationName                   = "insertWithOrganizationName"
	List                                         = "list"
	ListByCredentialId                           = "listByCredentialId"
	ListByOrganizationId                         = "listByOrganizationId"
	ListByOrganizationName                       = "listByOrganizationName"
	ListByTenantId                               = "listByTenantId"
	ListByTenantNameAndOrganizationName          = "listByTenantNameAndOrganizationName"
)

func New(path string, r *queryrepo.Repository) (*Registry, error) {
	if !strings.Contains(path, "?") {
		path = path + "?_pragma=foreign_keys(1)"
	} else if !strings.Contains(path, "&_pragma=foreign_keys(1)") {
		path = path + "&_pragma=pragma_foreign_keys(1)"
	}

	if r == nil {
		return nil, errors.New("repository is nil")
	}

	reg := &Registry{
		path:  path,
		repo:  r,
		stmts: make(map[string]*sql.Stmt),
	}
	return reg, nil
}

func NewWithDb(db *sql.DB, r *queryrepo.Repository) (*Registry, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}

	if r == nil {
		return nil, errors.New("repository is nil")
	}

	return &Registry{
		db:    db,
		repo:  r,
		stmts: make(map[string]*sql.Stmt),
	}, nil
}

type Registry struct {
	path  string
	db    *sql.DB
	repo  *queryrepo.Repository
	stmts map[string]*sql.Stmt
	mux   sync.Mutex
}

func (r *Registry) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

func (r *Registry) RawStmtExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	var err error
	var s *sql.Stmt
	if s, err = r.getRawStatementContext(ctx, query); err != nil {
		return nil, err
	}

	return s.ExecContext(ctx, args...)
}

func (r *Registry) Open() error {
	if r.db != nil {
		return nil
	}

	var err error
	if r.db, err = sql.Open("sqlite", r.path); err != nil {
		return err
	}
	return nil
}

func (r *Registry) RawStmtQueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	var err error
	var s *sql.Stmt
	if s, err = r.getRawStatementContext(ctx, query); err != nil {
		return nil, err
	}

	return s.QueryContext(ctx, args...)
}

func (r *Registry) getRawStatementContext(ctx context.Context, query string) (*sql.Stmt, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	k := strings.Join([]string{RawCollection, query}, "/")
	if stmt, ok := r.stmts[k]; ok {
		return stmt, nil
	}

	var err error
	if r.stmts[k], err = r.db.PrepareContext(ctx, query); err != nil {
		return nil, err
	}
	return r.stmts[k], nil
}

func (r *Registry) getStatement(ctx context.Context, collection, query string) (*sql.Stmt, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	k := strings.Join([]string{collection, query}, "/")
	if stmt, ok := r.stmts[k]; ok {
		return stmt, nil
	}

	var err error
	if r.stmts[k], err = r.repo.DbPrepareContext(ctx, r.db, collection, query); err != nil {
		return nil, err
	}
	return r.stmts[k], nil
}
