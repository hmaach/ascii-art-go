package asciiart

import "strings"

func GetInputAndBanner(args []string) (string, string) {
	var banner string
	var input string

	switch len(args) {
	case 2:
		banner = args[1] + ".txt"
		input = args[0]
	case 1:
		banner = "standard.txt"
		input = args[0]
	default:
		banner = "standard.txt"
		input = ""
	}

	return input, banner
}

// check if arguments are given as one string
func CheckSingleArgument(args []string) []string {
	if len(args) == 1 {
		tempArgs := strings.Fields(args[0])
		for _, v := range tempArgs {
			if strings.HasPrefix(v, "--") {
				args = tempArgs
			}
		}
		return args
	}
	return args
}
