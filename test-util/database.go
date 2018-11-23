package testutil

import (
	"fmt"
	"os"

	"github.com/spring2go/gravitee/log"

	"github.com/RichardKnop/go-fixtures"
	"github.com/jinzhu/gorm"
	"github.com/spring2go/gravitee/util/migrations"

	// Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateTestDatabase recreates the test database and
// runs migrations and fixtures as passed in, returning
// a pointer to the database
func CreateTestDatabase(dbPath string, migrationFunctions []func(*gorm.DB) error, fixtureFiles []string) (*gorm.DB, error) {

	// Init in-memory test database
	inMemoryDB, err := rebuildDatabase(dbPath)
	if err != nil {
		return nil, err
	}

	// Run all migrations
	migrations.MigrateAll(inMemoryDB, migrationFunctions)

	// Load data from data
	if err = fixtures.LoadFiles(fixtureFiles, inMemoryDB.DB(), "sqlite"); err != nil {
		return nil, err
	}

	return inMemoryDB, nil
}

// CreateTestDatabaseMysql is similar to CreateTestDatabase but it uses
// Mysql instead of sqlite, this is needed for testing packages that rely
// on some Mysql specifuc features (such as table inheritance)
func CreateTestDatabaseMysql(dbHost, dbUser, dbPass, dbName string, migrationFunctions []func(*gorm.DB) error, fixtureFiles []string) (*gorm.DB, error) {

	// Mysql test database
	db, err := rebuildDatabaseMysql(dbHost, dbUser, dbPass, dbName)
	if err != nil {
		return nil, err
	}

	// Run all migrations
	migrations.MigrateAll(db, migrationFunctions)

	// Load data from data
	if err = fixtures.LoadFiles(fixtureFiles, db.DB(), "mysql"); err != nil {
		return nil, err
	}

	return db, nil
}

// rebuildDatabase attempts to delete an existing in memory
// database and rebuild it, returning a pointer to it
func rebuildDatabase(dbPath string) (*gorm.DB, error) {
	// Delete the current database if it exists
	os.Remove(dbPath)

	// Init a new in-memory test database connection
	inMemoryDB, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return inMemoryDB, nil
}

// rebuildDatabase attempts to delete an existing Mysql
// database and rebuild it, returning a pointer to it
func rebuildDatabaseMysql(dbHost, dbUser, dbPass, dbName string) (*gorm.DB, error) {
	db, err := openMysqlDB(dbHost, dbUser, dbPass, dbName)
	if err != nil {
		return nil, err
	}

	if err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)).Error; err != nil {
		return nil, err
	}

	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error; err != nil {
		return nil, err
	}

	return openMysqlDB(dbHost, dbUser, dbPass, dbName)
}

func openMysqlDB(dbHost, dbUser, dbPass, dbName string) (*gorm.DB, error) {
	// Init a new mysql test database connection
	dbStr := fmt.Sprintf(
		"%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)
	log.INFO.Print(dbStr)
	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
