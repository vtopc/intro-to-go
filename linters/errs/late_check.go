package errs

import (
	"log"
)

func lateCheck() {
	resp, err := fnWithNilAndError()
	log.Println(resp.StatusCode)
	if err != nil {
		return
	}
}
