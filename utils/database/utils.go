package DatabaseUtils

import (
	StringUtils "github.com/josevi97/utils/string"
)

func ArrayInBrackets(array []string) string {
	return "(" + StringUtils.Join(array, ", ") + ")"
}
