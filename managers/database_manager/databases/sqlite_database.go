package databases

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const command = `
	CREATE TABLE IF NOT EXISTS command (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT,
		dir TEXT,
		command TEXT
	)
`

const command2Command = `
	CREATE TABLE IF NOT EXISTS command_to_command (
		command_id INTEGER PRIMARY KEY,
		command_id2 INTEGER PRIMARY KEY
	)
`

const template = `
	CREATE TABLE IF NOT EXISTS template (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT
	)
`

const command2Template = `
	CREATE TABLE IF NOT EXISTS command_to_command (
		command_id INTEGER PRIMARY KEY,
		template_id INTEGER PRIMARY KEY
	)
`

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
	db.conn.Exec(command)
	db.conn.Exec(command2Command)
	db.conn.Exec(template)
	db.conn.Exec(command2Template)
}

func (db *Sqlite) GetPath() string {
	return db.pathname
}

func NewSqlite(pathname string) *Sqlite {
	return &Sqlite{
		pathname: pathname,
	}
}
