package errs

func staticcheck() (err error) {
	err = fnWithError()
	err = fnNoError()
	return
}
