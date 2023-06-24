package databases

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	DatabaseUtils "github.com/josevi97/utils/database"
	StringUtils "github.com/josevi97/utils/string"
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
	keys, values := StringUtils.GetDataFromMap(data)
	array := StringUtils.Map(values, "?")

	var builder strings.Builder

	// TODO: find a way to make a function chain
	builder.WriteString("INSERT INTO ")
	builder.WriteString(table + " ")
	builder.WriteString(DatabaseUtils.ArrayInBrackets(keys) + " ")
	builder.WriteString("VALUES ")
	builder.WriteString(DatabaseUtils.ArrayInBrackets(array))

	query := builder.String()

	Log.Info("query to create a new instance: " + query)

	stm, _ := db.conn.Prepare(query)
	stm.Exec(values)
}

func NewSqlite(pathname string) *Sqlite {
	return &Sqlite{
		pathname: pathname,
	}
}
