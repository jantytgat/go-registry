package registry

import (
	"context"
	"database/sql"
	"fmt"
)

type ConnectionType struct {
	Id   int64
	Name string
}

func (r *Registry) DeleteConnectionTypeById(ctx context.Context, id int64) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, ConnectionTypesCollection, DeleteById); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(id); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) DeleteConnectionTypeByName(ctx context.Context, name string) (int64, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, ConnectionTypesCollection, DeleteByName); err != nil {
		return 0, err
	}

	var result sql.Result
	if result, err = s.Exec(name); err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *Registry) GetConnectionTypeById(ctx context.Context, id int64) (ConnectionType, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, ConnectionTypesCollection, GetById); err != nil {
		return ConnectionType{}, err
	}

	var o ConnectionType
	return o, s.QueryRow(id).Scan(&o.Id, &o.Name)
}

func (r *Registry) GetConnectionTypeByName(ctx context.Context, name string) (ConnectionType, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, ConnectionTypesCollection, GetByName); err != nil {
		return ConnectionType{}, err
	}

	var o ConnectionType
	return o, s.QueryRow(name).Scan(&o.Id, &o.Name)
}

func (r *Registry) InsertConnectionType(ctx context.Context, guid, name string) (ConnectionType, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, ConnectionTypesCollection, Insert); err != nil {
		return ConnectionType{}, err
	}

	var result sql.Result
	if result, err = s.Exec(guid, name); err != nil {
		return ConnectionType{}, err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return ConnectionType{}, err
	}
	return ConnectionType{
		Id:   id,
		Name: name,
	}, err
}

func (r *Registry) ListConnectionTypes(ctx context.Context) ([]ConnectionType, error) {
	var err error
	var s *sql.Stmt

	if s, err = r.getStatement(ctx, ConnectionTypesCollection, List); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = s.Query(); err != nil {
		fmt.Println("Error querying listing connection_types")
		return nil, err
	}
	defer rows.Close()

	var connection_types []ConnectionType
	for rows.Next() {
		var o ConnectionType
		if err = rows.Scan(&o.Id, &o.Name); err != nil {
			fmt.Println("Error scanning listing connection_types")
			return nil, err
		}
		connection_types = append(connection_types, o)
	}
	return connection_types, nil
}
