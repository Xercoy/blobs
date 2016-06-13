package glob

func validUnit(unit string) bool {
	switch unit {
	case "B":
		return true
	case "MB":
		return true
	case "GB":
		return true
	case "TB":
		return true
	}
	return false
}

func validMode(source string) bool {
	switch source {
	case "default":
		return true
	case "urandom":
		return true
	case "random":
		return true
	}

	return false
}
