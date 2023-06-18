package error

type ErrorType int

const (
	NONE = iota
	FILESYSTEM_DIRECTRY_ALREADY_EXISTS
	FILESYSTEM_CREATE_DIRECTORY
)

type Error struct {
	ErrorType ErrorType
}
