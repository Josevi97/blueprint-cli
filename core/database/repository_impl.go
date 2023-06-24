package database

type RepositoryImpl struct {
	db        Database
	tablename string
}

func (repository *RepositoryImpl) whileConnection(callback func()) {
	repository.db.Open()
	callback()
	repository.db.Close()
}

func (repository *RepositoryImpl) Create(data map[string]string) {
	callback := func() {
		repository.db.Create(repository.tablename, data)
	}

	repository.whileConnection(callback)
}

func NewRepository(db Database, tablename string) Repository {
	return &RepositoryImpl{
		db:        db,
		tablename: tablename,
	}
}
