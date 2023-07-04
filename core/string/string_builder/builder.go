package string_builder

type Builder interface {
	// used to write some string in the current builder
	Write(s string) Builder

	// returns the final string
	Build() string
}
