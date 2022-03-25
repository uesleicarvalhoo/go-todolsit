package repository

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(&entity.Task{})
}

func DBMigrate(dbInstance *gorm.DB, dbName string) error {
	db, err := dbInstance.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "Failed to instantiate postgres driver")
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://internal/repository/migrations", dbName, driver)
	if err != nil {
		return errors.Wrap(err, "Failed to create migrate instance")
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "Faile to apply migrations up")
	}
	return nil
}
