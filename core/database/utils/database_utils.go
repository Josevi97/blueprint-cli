package DatabaseUtils

import (
	StringUtils "github.com/josevi97/core/string/utils"
)

func ArrayInBrackets(array []string) string {
	return "(" + StringUtils.Join(array, ", ") + ")"
}
