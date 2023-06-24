package Database

type Repository interface {
	// SQL insert interace
	Create(data map[string]string)
}
