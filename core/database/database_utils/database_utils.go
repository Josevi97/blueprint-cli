package database_utils

import "github.com/josevi97/core/string/string_utils"

func ArrayInBrackets(array []string) string {
	return "(" + string_utils.Join(array, ", ") + ")"
}
