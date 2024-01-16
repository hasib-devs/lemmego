package framework

import (
	"errors"
	"fmt"
	"log"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var (
	ErrUnknownDialect = errors.New("unknown dialect")
	ErrNoSuchDatabase = errors.New("no such database exists")
)

type DBSession interface {
	db.Session
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func DBExists(dialect string, dbConfig *DBConfig) (bool, error) {
	var sess db.Session
	var err error

	defer func() {
		if sess != nil {
			log.Println("closing session ", sess.Name())
			sess.Close()
		}
	}()

	if dialect == "postgres" {
		settings := postgresql.ConnectionURL{
			Host:     dbConfig.Host,
			User:     dbConfig.User,
			Password: dbConfig.Password,
		}

		sess, err = postgresql.Open(settings)
		if err != nil {
			return false, fmt.Errorf("failed to connect: %w", err)
		}

		if row, err := sess.SQL().QueryRow("SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('" + dbConfig.Database + "')"); err != nil {
			return false, fmt.Errorf("failed to check db existence: %w", err)
		} else {
			var database string
			row.Scan(&database)
			if database != dbConfig.Database {
				return false, ErrNoSuchDatabase
			}
			return true, nil
		}
	}

	if dialect == "mysql" {
		settings := mysql.ConnectionURL{
			Host:     dbConfig.Host,
			User:     dbConfig.User,
			Password: dbConfig.Password,
		}

		sess, err = mysql.Open(settings)
		if err != nil {
			return false, fmt.Errorf("failed to connect: %w", err)
		}

		if row, err := sess.SQL().QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '" + dbConfig.Database + "'"); err != nil {
			return false, fmt.Errorf("failed to check db existence: %w", err)
		} else {
			var database string
			row.Scan(&database)
			if database != dbConfig.Database {
				return false, ErrNoSuchDatabase
			}
			return true, nil
		}
	}

	return false, ErrUnknownDialect
}

func CreateDB(dialect string, dbConfig *DBConfig) error {
	var sess db.Session
	var err error

	defer func() {
		if sess != nil {
			log.Println("closing session ", sess.Name())
			sess.Close()
		}
	}()

	if sess, err = CreateSession(dialect, dbConfig, true); err != nil {
		return err
	}

	if dialect == "postgres" {
		res, err := sess.SQL().Exec("CREATE DATABASE " + dbConfig.Database + " WITH OWNER " + dbConfig.User)
		if err != nil {
			return err
		}
		if res != nil {
			return nil
		}
	}

	if dialect == "mysql" {
		res, err := sess.SQL().Exec("CREATE DATABASE IF NOT EXISTS " + dbConfig.Database)
		if err != nil {
			return err
		}
		if res != nil {
			return nil
		}
	}

	return nil
}

func CreateSession(dialect string, dbConfig *DBConfig, ignoreDb bool) (DBSession, error) {
	if dialect == "postgres" {
		settings := postgresql.ConnectionURL{
			Host:     dbConfig.Host,
			User:     dbConfig.User,
			Password: dbConfig.Password,
		}

		if !ignoreDb {
			settings.Database = dbConfig.Database
		}

		sess, err := postgresql.Open(settings)
		if err != nil {
			return nil, fmt.Errorf("failed to connect: %w", err)
		}

		return sess, nil
	}

	if dialect == "mysql" {
		settings := mysql.ConnectionURL{
			Host:     dbConfig.Host,
			User:     dbConfig.User,
			Password: dbConfig.Password,
		}

		if !ignoreDb {
			settings.Database = dbConfig.Database
		}

		sess, err := mysql.Open(settings)
		if err != nil {
			return nil, fmt.Errorf("failed to connect: %w", err)
		}

		return sess, nil
	}

	return nil, ErrUnknownDialect
}

func ConnectDB(dialect string, dbConfig *DBConfig) (DBSession, error) {
	if sess, err := CreateSession(dialect, dbConfig, true); err != nil {
		return nil, err
	} else {
		if exists, _ := DBExists(dialect, dbConfig); !exists {
			if err := CreateDB(dialect, dbConfig); err != nil {
				return nil, err
			}
		} else {
			sess.Close()
			sess, err = CreateSession(dialect, dbConfig, false)
			if err != nil {
				return nil, err
			}
		}
		return sess, nil
	}
}
