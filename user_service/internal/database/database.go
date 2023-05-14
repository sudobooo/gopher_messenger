package database

import (
	"database/sql"
)

type Database struct {
	db     *sql.DB
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (d *Database) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}

	return nil
}
