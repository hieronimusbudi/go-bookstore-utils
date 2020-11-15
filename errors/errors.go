package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//RestErr interface
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
	Dummy() string
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v", e.ErrMessage, e.ErrStatus, e.ErrMessage, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func (e restErr) Dummy() string {
	return e.ErrMessage
}

//NewRestError func
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

//NewRestErrorFromBytes func
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

//NewBadRequestError func
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError func
func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

//NewUnauthorizedError func
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

//NewInternalServerError func
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
