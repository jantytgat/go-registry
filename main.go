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

	if err = initializeLocal(); err != nil {
		log.Fatal(err)
	}

	if err = migrateDb(); err != nil {
		log.Fatal(err)
	}

	// Create query repository from embedded files
	if rsql, err = queryrepo.NewFromFs(queriesFs, "assets/queries"); err != nil {
		log.Fatal(err)
	}

	var organizations []registry.Organization
	if organizations, err = registry.ListOrganizations(rsql, db); err != nil {
		log.Fatal(err)
	}

	for _, o := range organizations {
		fmt.Println(o)
	}
}

func initializeInMemory() error {
	var err error
	db, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		fmt.Println("Error opening database")
		return err
	}
	return nil
}

func initializeLocal() error {
	var err error
	db, err = sql.Open("sqlite", "./registry.db")
	if err != nil {
		fmt.Println("Error opening database")
		return err
	}
	return nil
}

func migrateDb() error {
	var err error
	var src source.Driver
	if src, err = iofs.New(migrationFs, "assets/migrations"); err != nil {
		fmt.Println("Error opening migrations source")
		return err
	}

	var driver database.Driver
	if driver, err = sqlite.WithInstance(db, &sqlite.Config{}); err != nil {
		fmt.Println("Error opening migrations destination")
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
