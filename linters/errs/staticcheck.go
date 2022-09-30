package errs

func staticcheck1() (err error) {
	err = fnWithError()
	err = fnNoError()
	return
}

func staticcheck2() (err error) {
	err = fnNoError()
	return fnWithError()
}
