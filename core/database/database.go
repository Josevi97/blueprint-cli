package database

type Database interface {
	// opens the database connection
	Open(path string) bool

	// closes the database connection
	Close()

	// returns the pathname
	GetPath() string

	// executes every sql query in order to create the database structure
	DoMigrations()
}
