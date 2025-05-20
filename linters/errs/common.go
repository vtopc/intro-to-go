package errs

import (
	"errors"
	"net/http"
)

func fnWithError() error {
	return errors.New("some error")
}

func fnNoError() error {
	return nil
}

func fnWithNilAndError() (*http.Response, error) {
	return nil, errors.New("some error")
}
