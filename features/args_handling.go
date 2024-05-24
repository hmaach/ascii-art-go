package asciiart

func GetInputAndBanner(args []string) (string, string) {
	var banner string = "standard.txt"
	var input string

	switch len(args) {
	case 2:
		banner = args[1]
		input = args[0]
	case 1:
		input = args[0]
	default:
		input = ""
	}

	return input, banner
}
