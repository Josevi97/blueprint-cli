package Database

type Database interface {
	// opens the database connection
	Open() bool

	// closes the database connection
	Close()

	// executes every sql query in order to create the database structure
	Migrate()

	// returns the pathname
	GetPath() string

	// SQL Queries

	// Get()

	// GetAll()

	Create(table string, data map[string]string)

	// Update()

	// Delete()
}
