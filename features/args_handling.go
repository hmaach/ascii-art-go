package asciiart

import "strings"

func GetInputAndBanner(args []string) (string, string) {
	var banner string
	var input string

	if len(args) == 2 {
		banner = args[1]
		banner += ".txt"
	} else {
		banner = "standard.txt"
	}

	if len(args) >= 1 {
		input = args[0]
	} else {
		input = ""
	}
	return input, banner
}

// check if arguments are given as one string
func CheckSingleArgument(args []string) []string {
	tempArgs := strings.Fields(args[0])
	if len(args) == 1 {
		for _, v := range tempArgs {
			if strings.HasPrefix(v, "--") {
				args = tempArgs
			}
		}
		return tempArgs
	}
	return args
}
