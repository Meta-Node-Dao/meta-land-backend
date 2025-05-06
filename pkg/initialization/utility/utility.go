package utility

// Init all the utility package which need init at startup
func Init() (err error) {
	err = initSequence()
	if err != nil {
		return
	}

	return
}

func ConvertToInterfaceSlice[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
