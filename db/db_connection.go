package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DbConnection struct {
	sqlDb *sqlx.DB
}

type Args map[string]interface{}

var (
	DbConn *DbConnection
)

func ConnectDb() error {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres host=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(0)

	DbConn = &DbConnection{sqlDb: db}
	return nil
}

func (c *DbConnection) ExecuteSelect(dest interface{}, query string, args Args) error {

	stmt, err := c.sqlDb.PrepareNamed(query)

	if err != nil {
		return errors.Wrapf(err, "error on prepare")
	}

	err = stmt.Select(dest, args)
	if err != nil {
		return errors.Wrapf(err, "error on get")
	}
	return nil
}

func (c *DbConnection) ExecuteSelectWithArgs(dest interface{}, query string, keyvals ...interface{}) error {

	args := make(map[string]interface{}, len(keyvals)/2)
	for i := 0; i < len(keyvals); i += 2 {
		k, v := keyvals[i], keyvals[i+1]
		args[fmt.Sprint(k)] = v
	}

	stmt, err := c.sqlDb.PrepareNamed(query)

	if err != nil {
		return errors.Wrapf(err, "error on prepare")
	}

	err = stmt.Select(dest, args)
	if err != nil {
		return errors.Wrapf(err, "error on get")
	}
	return nil
}

func (c *DbConnection) ExecuteUpdate(query string, args Args) error {
	stmt, err := c.sqlDb.PrepareNamed(query)

	if err != nil {
		return errors.Wrapf(err, "error on prepare")
	}

	_, err = stmt.Exec(args)
	if err != nil {
		return errors.Wrapf(err, "error on exec")
	}
	return nil
}

func (c *DbConnection) Close() {
	c.sqlDb.Close()
}
