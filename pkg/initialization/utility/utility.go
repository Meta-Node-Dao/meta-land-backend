package utility

// Init all the utility package which need init at startup
func Init() (err error) {
	err = initSequence()
	if err != nil {
		return
	}

	return
}
