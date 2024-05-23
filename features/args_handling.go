package asciiart

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
