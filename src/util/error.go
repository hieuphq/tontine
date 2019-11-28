package util

import (
	"net/http"

	"github.com/hieuphq/tontine/src/model/errors"
)

// ParseErrorCode parse error code from errors.Error
func ParseErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch arg := err.(type) {
	case *errors.Error:
		return int(arg.Code)

	case error:
		return http.StatusInternalServerError

	default:
		return http.StatusOK
	}
}
