package database

import (
	"errors"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // comento como si supiera de lo que hablo
)

// NewDatabase ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}

/*
// Team tiene un nombre
type Team struct {
	ID   int    `db:"id"`
	name string `db:"name"`
}

	db.MustExec("INSERT INTO team (name) VALUES (?)", "jane doe")
	rows, err := db.Queryx("SELECT id, name FROM team")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var t Team
		rows.StructScan(&t)
		fmt.Println(t)
	}
*/
