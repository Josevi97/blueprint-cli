package ArgsUtils

func GetNextArgs(step int, args []string) []string {
	if args == nil || len(args) < step {
		return []string{}
	}

	return args[step:]
}
