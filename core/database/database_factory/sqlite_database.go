package database_factory

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/josevi97/core/database/database_utils"
	"github.com/josevi97/core/string/string_builder"
	"github.com/josevi97/core/string/string_utils"
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

func (db *Sqlite) Create(table string, data map[string]string) {
	keys, values := string_utils.GetDataFromMap(data)
	query := string_builder.NewBuilder().
		Write("INSERT INTO ").
		Write(table + " ").
		Write(database_utils.ArrayInBrackets(keys) + " ").
		Write("VALUES ").
		Write(database_utils.ArrayInBrackets(string_utils.Map(values, "?"))).
		Build()

	Log.Info("query to create a new instance: " + query)

	stm, _ := db.conn.Prepare(query)
	stm.Exec(values)
}

func NewSqlite(pathname string) *Sqlite {
	return &Sqlite{
		pathname: pathname,
	}
}
