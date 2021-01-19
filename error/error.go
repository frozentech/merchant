package error

import (
	"fmt"
	"net/http"

	"github.com/frozentech/merchant/model"
)

// Error codes
const (
	NoError = iota
	RequestBodyEmpty
	RequestBodyStatusUnprocessableEntity
	RecordNotFound
	DatabaseOperationFailure
	DuplicateEmailAddress
	UploadFailed
	InvalidEmailAddress
	InvalidPageValue
	InvalidLimitValue
)

// ErrorMessages ...
var ErrorMessages = map[int]model.Error{
	NoError: model.Error{
		Status:  200,
		Code:    fmt.Sprintf("%04d", NoError),
		Message: http.StatusText(http.StatusOK),
	},
	RequestBodyEmpty: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", RequestBodyEmpty),
		Message: "Empty Request Body",
	},
	RequestBodyStatusUnprocessableEntity: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", RequestBodyStatusUnprocessableEntity),
		Message: http.StatusText(http.StatusUnprocessableEntity),
	},
	RecordNotFound: model.Error{
		Status:  http.StatusNotFound,
		Code:    fmt.Sprintf("%04d", RecordNotFound),
		Message: http.StatusText(http.StatusNotFound),
	},
	DatabaseOperationFailure: model.Error{
		Status:  http.StatusServiceUnavailable,
		Code:    fmt.Sprintf("%04d", DatabaseOperationFailure),
		Message: "Database Operation Failure",
	},
	UploadFailed: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", UploadFailed),
		Message: "Upload Failed",
	},
	InvalidPageValue: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", InvalidPageValue),
		Message: "Invalid Page Value",
	},
	InvalidLimitValue: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", InvalidLimitValue),
		Message: "Invalid Limit Value",
	},
	DuplicateEmailAddress: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", DuplicateEmailAddress),
		Message: "Duplicate Email Address",
	},
	InvalidEmailAddress: model.Error{
		Status:  http.StatusBadRequest,
		Code:    fmt.Sprintf("%04d", InvalidEmailAddress),
		Message: "Invalid Email Address",
	},
}

// StatusRecord ...
func StatusRecord(code int) model.Response {

	err := ErrorMessages[code]

	if code == NoError {
		return model.Response{
			Success: true,
		}
	}

	return model.Response{
		Success: false,
		Error:   &err,
	}

}
