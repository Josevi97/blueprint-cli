package databases

import (
	"database/sql"

	FileSystemUtils "github.com/josevi97/utils/file_system"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	pathname string
	conn     *sql.DB
}

func (db *Sqlite) Open(path string) bool {
	FileSystemUtils.Join(db.pathname)

	conn, err := sql.Open("sqlite3", db.GetPath())
	db.conn = conn

	return err == nil
}

func (db *Sqlite) Close() {
	db.conn.Close()
}

func (db *Sqlite) GetPath() string {
	return db.pathname
}

func (db *Sqlite) DoMigrations() {
	db.conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`)
}

func NewSqlite(pathname string) *Sqlite {
	return &Sqlite{
		pathname: pathname,
	}
}
