package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"

	"github.com/jantytgat/go-registry/pkg/registry"
)

//go:embed assets/migrations/*
var migrationFs embed.FS

//go:embed assets/queries/*
var queriesFs embed.FS

var rsql *queryrepo.Repository

func main() {
	var err error
	var path = "./registry.db"
	// var path = ":memory:"

	if err = migrateDb(path); err != nil {
		fmt.Println("Failed to migrate db")
		panic(err)
	}

	// Create query repository from embedded files
	if rsql, err = queryrepo.NewFromFs(queriesFs, "assets/queries"); err != nil {
		fmt.Println("Failed to load queries")
		panic(err)
	}

	var r *registry.Registry
	if r, err = registry.New(path, rsql); err != nil {
		fmt.Println("Failed to create registry")
		panic(err)
	}
	if err = r.Open(); err != nil {
		fmt.Println("Failed to open registry")
		panic(err)
	}
	defer r.Close()

	ctx := context.Background()

	var organization registry.Organization
	if organization, err = r.GetOrganizationByGuid(ctx, "0"); err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Failed to get organization", "0")
		panic(err)
	}
	fmt.Println(organization)

	var organizations []registry.Organization
	if organizations, err = r.ListOrganizations(ctx); err != nil {
		fmt.Println("Failed to list organizations")
		panic(err)
	}

	// var rowsAffected int64
	// if rowsAffected, err = r.DeleteOrganizationByName(ctx, "default"); err != nil {
	// 	fmt.Println("Failed to delete organization")
	// 	panic(err)
	// }
	// fmt.Println("Delete default org", rowsAffected, err)
	//
	fmt.Println("Listing organizations:")
	for _, o := range organizations {
		fmt.Println(o)
	}

	var tenants []registry.Tenant
	if tenants, err = r.ListTenants(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listing tenants:")
	for _, t := range tenants {
		fmt.Println(t)
	}

}

func migrateDb(path string) error {
	var err error

	var db *sql.DB
	if db, err = sql.Open("sqlite", path); err != nil {
		return err
	}
	defer db.Close()

	var src source.Driver
	if src, err = iofs.New(migrationFs, "assets/migrations"); err != nil {
		return err
	}

	var driver database.Driver
	if driver, err = sqlite.WithInstance(db, &sqlite.Config{}); err != nil {
		return err
	}

	var m *migrate.Migrate
	if m, err = migrate.NewWithInstance("fs", src, "sqlite", driver); err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
