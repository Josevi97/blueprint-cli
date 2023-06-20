package ErrorUtils

import (
	"github.com/josevi97/core/error"
)

func IsError(err uint) bool {
	return err != error.NONE
}

func Throw(v any) {
	panic(v)
}

func AssertError(err uint) {
	if IsError(err) {
		Throw(err)
	}
}
