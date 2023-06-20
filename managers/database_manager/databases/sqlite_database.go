package databases

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	pathname string
	conn     *sql.DB
}

func (db *Sqlite) Open() bool {
	conn, err := sql.Open("sqlite3", db.GetPath())
	db.conn = conn

	return err == nil
}

func (db *Sqlite) Close() {
	db.conn.Close()
}

func (db *Sqlite) Migrate() {
	db.conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`)
}

func (db *Sqlite) GetPath() string {
	return db.pathname
}

func NewSqlite(pathname string) *Sqlite {
	return &Sqlite{
		pathname: pathname,
	}
}
