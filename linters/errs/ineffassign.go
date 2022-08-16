package errs

func ineffassign() error {
	err := fnWithError()
	err = fnNoError()
	if err != nil {
		return err
	}

	return nil
}
