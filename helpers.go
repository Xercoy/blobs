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

// Maybe return error?
func BytesInUnit(unit string) int {
	var byteAmount int
	switch unit {
	case "B":
		byteAmount = 1
	case "KB":
		byteAmount = 1024
	case "MB":
		byteAmount = 1048576
	case "GB":
		byteAmount = 1073741824
	case "TB":
		byteAmount = 1099511627776
	}

	return byteAmount
}
