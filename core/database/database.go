package database

type Database interface {
	// opens the database connection
	Open() bool

	// closes the database connection
	Close()

	// executes every sql query in order to create the database structure
	Migrate()

	// returns the pathname
	GetPath() string
}
