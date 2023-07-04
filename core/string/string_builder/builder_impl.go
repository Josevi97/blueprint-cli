package string_builder

import "strings"

type BuilderImpl struct {
	builder strings.Builder
}

func (builder *BuilderImpl) Write(s string) Builder {
	builder.builder.WriteString(s)
	return builder
}

func (builder *BuilderImpl) Build() string {
	return builder.builder.String()
}

func NewBuilder() Builder {
	var builder strings.Builder

	return &BuilderImpl{
		builder: builder,
	}
}
