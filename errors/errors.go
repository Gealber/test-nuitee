package errors

import "net/http"

// ErrService error containing information about the htttp status
// code to use depending on the error and a message of the error
// Meant to be used on controller package or any other place where
// assigning an http code to a given error is required
type ErrService struct {
	code int
	msg  string
}

func NewErrService(code int, msg string) ErrService {
	return ErrService{code: code, msg: msg}
}

func (e ErrService) Error() string {
	return e.msg
}

func (e ErrService) Code() int {
	return e.code
}

// ParseServiceError parses a given error extracting its code and message
// in case the error is not of type ErrService, we return an 500 with the error message
func ParseServiceError(err error) (int, error) {
	switch errT := err.(type) {
	case ErrService:
		return errT.code, errT
	default:
		return http.StatusInternalServerError, err
	}
}
