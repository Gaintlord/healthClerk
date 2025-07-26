package database

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db gorm.DB
}

func (d *Database) Connect() error {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("unable to connect to database", slog.String("error", err.Error()))
		return err
	}

	d.Db = *db
	slog.Info("### DataBase Connection establised ###")
	return nil
}

func (d *Database) GetDB() gorm.DB {
	return d.Db
}
