package errs

import (
	"errors"
)

func fnWithError() error {
	return errors.New("some error")
}

func fnNoError() error {
	return nil
}
