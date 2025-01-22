package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Driver for sqlite3
)

type DB struct {
    conn *sql.DB
}

func (db *DB) Conn() error {
    var err error
    db.conn, err = sql.Open("sqlite3", "password.db")

    if err != nil {
        return err
    }

    return nil
}

func (db *DB) Init() error {
    tableStmt := `CREATE TABLE password(
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL
    )`

    if _, err := db.conn.Exec(tableStmt); err != nil {
        return err
    }

    return nil
}
