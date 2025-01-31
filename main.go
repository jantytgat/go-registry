package main

import (
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

var db *sql.DB
var rsql *queryrepo.Repository

func main() {
	var err error

	if err = initializeInMemory(); err != nil {
		log.Fatal(err)
	}

	if err = migrateDb(); err != nil {
		log.Fatal(err)
	}

	// Create query repository from embedded files
	if rsql, err = queryrepo.NewFromFs(queriesFs, "assets/queries"); err != nil {
		log.Fatal(err)
	}

	var organization registry.Organization
	if organization, err = registry.GetOrganizationByGuid(rsql, db, "0"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(organization)

	var organization2 registry.Organization
	if organization2, err = registry.AddOrganization(rsql, db, "1", "corelayer"); err == nil {
		fmt.Println(organization2)
	}

	var organizations []registry.Organization
	if organizations, err = registry.ListOrganizations(rsql, db); err != nil {
		log.Fatal(err)
	}

	var rowsAffected int64
	rowsAffected, err = registry.DeleteOrganizationByName(rsql, db, "default")
	fmt.Println("Delete default org", rowsAffected, err)

	fmt.Println("Listing organizations:")
	for _, o := range organizations {
		fmt.Println(o)
	}

	var tenants []registry.Tenant
	if tenants, err = registry.ListTenants(rsql, db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listing tenants:")
	for _, t := range tenants {
		fmt.Println(t)
	}

}

func initializeInMemory() error {
	var err error
	db, err = registry.Connect(":memory:")
	if err != nil {
		return err
	}
	return nil
}

func initializeLocal() error {
	var err error
	db, err = registry.Connect("./registry.db")
	if err != nil {
		return err
	}
	return nil
}

func migrateDb() error {
	var err error
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
