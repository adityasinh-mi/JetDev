package utils

import (
	"errors"
	"fmt"
	"strings"
)

/**
** EVERY TIME YOU ADD A NEW ERROR CODE CONSTANTS PLEASE ALSO ADD THE CONSTANTS IN Error_test.go and make sure the test passed
**/
const (
	// http code below
	ERROR_ACCEPTED                        = 202
	ERROR_BAD_GATEWAY                     = 502
	ERROR_BAD_REQUEST                     = 400
	ERROR_CONFLICT                        = 409
	ERROR_CONTINUE                        = 100
	ERROR_CREATED                         = 201
	ERROR_EXPECTATION_FAILED              = 417
	ERROR_FAILED_DEPENDENCY               = 424
	ERROR_FORBIDDEN                       = 403
	ERROR_GATEWAY_TIMEOUT                 = 504
	ERROR_GONE                            = 410
	ERROR_HTTP_VERSION_NOT_SUPPORTED      = 505
	ERROR_IM_A_TEAPOT                     = 418
	ERROR_INSUFFICIENT_SPACE_ON_RESOURCE  = 419
	ERROR_INSUFFICIENT_STORAGE            = 507
	ERROR_INTERNAL_SERVER_ERROR           = 500
	ERROR_LENGTH_REQUIRED                 = 411
	ERROR_LOCKED                          = 423
	ERROR_METHOD_FAILURE                  = 420
	ERROR_METHOD_NOT_ALLOWED              = 405
	ERROR_MOVED_PERMANENTLY               = 301
	ERROR_MOVED_TEMPORARILY               = 302
	ERROR_MULTI_STATUS                    = 207
	ERROR_MULTIPLE_CHOICES                = 300
	ERROR_NETWORK_AUTHENTICATION_REQUIRED = 511
	ERROR_NO_CONTENT                      = 204
	ERROR_NON_AUTHORITATIVE_INFORMATION   = 203
	ERROR_NOT_ACCEPTABLE                  = 406
	ERROR_NOT_FOUND                       = 404
	ERROR_NOT_IMPLEMENTED                 = 501
	ERROR_NOT_MODIFIED                    = 304
	ERROR_OK                              = 200
	ERROR_PARTIAL_CONTENT                 = 206
	ERROR_PAYMENT_REQUIRED                = 402
	ERROR_PERMANENT_REDIRECT              = 308
	ERROR_PRECONDITION_FAILED             = 412
	ERROR_PRECONDITION_REQUIRED           = 428
	ERROR_PROCESSING                      = 102
	ERROR_PROXY_AUTHENTICATION_REQUIRED   = 407
	ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE = 431
	ERROR_REQUEST_TIMEOUT                 = 408
	ERROR_REQUEST_TOO_LONG                = 413
	ERROR_REQUEST_URI_TOO_LONG            = 414
	ERROR_REQUESTED_RANGE_NOT_SATISFIABLE = 416
	ERROR_RESET_CONTENT                   = 205
	ERROR_SEE_OTHER                       = 303
	ERROR_SERVICE_UNAVAILABLE             = 503
	ERROR_SWITCHING_PROTOCOLS             = 101
	ERROR_TEMPORARY_REDIRECT              = 307
	ERROR_TOO_MANY_REQUESTS               = 429
	ERROR_UNAUTHORIZED                    = 401
	ERROR_UNPROCESSABLE_ENTITY            = 422
	ERROR_UNSUPPORTED_MEDIA_TYPE          = 415
	ERROR_USE_PROXY                       = 305
)

const (
	ERROR_FIELD_REQUIRED = iota + 1101
)

var messages = map[int]string{
	// http code message below
	ERROR_ACCEPTED:                        "Accepted",
	ERROR_BAD_GATEWAY:                     "Bad Gateway",
	ERROR_BAD_REQUEST:                     "Bad Request",
	ERROR_CONFLICT:                        "Conflict",
	ERROR_CONTINUE:                        "Continue",
	ERROR_CREATED:                         "Created",
	ERROR_EXPECTATION_FAILED:              "Expectation Failed",
	ERROR_FAILED_DEPENDENCY:               "Failed Dependency",
	ERROR_FORBIDDEN:                       "Forbidden",
	ERROR_GATEWAY_TIMEOUT:                 "Gateway Timeout",
	ERROR_GONE:                            "Gone",
	ERROR_HTTP_VERSION_NOT_SUPPORTED:      "HTTP Version Not Supported",
	ERROR_IM_A_TEAPOT:                     "I'm a teapot",
	ERROR_INSUFFICIENT_SPACE_ON_RESOURCE:  "Insufficient Space on Resource",
	ERROR_INSUFFICIENT_STORAGE:            "Insufficient Storage",
	ERROR_INTERNAL_SERVER_ERROR:           "Server Error",
	ERROR_LENGTH_REQUIRED:                 "Length Required",
	ERROR_LOCKED:                          "Locked",
	ERROR_METHOD_FAILURE:                  "Method Failure",
	ERROR_METHOD_NOT_ALLOWED:              "Method Not Allowed",
	ERROR_MOVED_PERMANENTLY:               "Moved Permanently",
	ERROR_MOVED_TEMPORARILY:               "Moved Temporarily",
	ERROR_MULTI_STATUS:                    "Multi-Status",
	ERROR_MULTIPLE_CHOICES:                "Multiple Choices",
	ERROR_NETWORK_AUTHENTICATION_REQUIRED: "Network Authentication Required",
	ERROR_NO_CONTENT:                      "No Content",
	ERROR_NON_AUTHORITATIVE_INFORMATION:   "Non Authoritative Information",
	ERROR_NOT_ACCEPTABLE:                  "Not Acceptable",
	ERROR_NOT_FOUND:                       "Not Found",
	ERROR_NOT_IMPLEMENTED:                 "Not Implemented",
	ERROR_NOT_MODIFIED:                    "Not Modified",
	ERROR_OK:                              "OK",
	ERROR_PARTIAL_CONTENT:                 "Partial Content",
	ERROR_PAYMENT_REQUIRED:                "Payment Required",
	ERROR_PERMANENT_REDIRECT:              "Permanent Redirect",
	ERROR_PRECONDITION_FAILED:             "Precondition Failed",
	ERROR_PRECONDITION_REQUIRED:           "Precondition Required",
	ERROR_PROCESSING:                      "Processing",
	ERROR_PROXY_AUTHENTICATION_REQUIRED:   "Proxy Authentication Required",
	ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE: "Request Header Fields Too Large",
	ERROR_REQUEST_TIMEOUT:                 "Request Timeout",
	ERROR_REQUEST_TOO_LONG:                "Request Entity Too Large",
	ERROR_REQUEST_URI_TOO_LONG:            "Request-URI Too Long",
	ERROR_REQUESTED_RANGE_NOT_SATISFIABLE: "Requested Range Not Satisfiable",
	ERROR_RESET_CONTENT:                   "Reset Content",
	ERROR_SEE_OTHER:                       "See Other",
	ERROR_SERVICE_UNAVAILABLE:             "Service Unavailable",
	ERROR_SWITCHING_PROTOCOLS:             "Switching Protocols",
	ERROR_TEMPORARY_REDIRECT:              "Temporary Redirect",
	ERROR_TOO_MANY_REQUESTS:               "Too Many Requests",
	ERROR_UNAUTHORIZED:                    "Unauthorized",
	ERROR_UNPROCESSABLE_ENTITY:            "Unprocessable Entity",
	ERROR_UNSUPPORTED_MEDIA_TYPE:          "Unsupported Media Type",
	ERROR_USE_PROXY:                       "Use Proxy",
}

var Valmessages = map[string]int{
	// custom error message below
	"ERROR_FIELD_REQUIRED": 1101,
}

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

func ThrowError(code int) error {
	err := APIError{
		Code:    code,
		Message: "ERROR",
	}
	if msg, ok := messages[code]; ok {
		err.Message = msg
	}
	return &err
}

func GetErrorCode(str string, valType string) int {
	errStr := "ERROR_" + strings.ToUpper(str) + "_" + strings.ToUpper(valType)
	if Valmessages[errStr] > 0 {
		return Valmessages[errStr]
	}
	return ERROR_FIELD_REQUIRED

}
func ThrowBadRequest(code int) error {
	err := APIError{
		Code:    400,
		Message: "ERROR",
	}
	if msg, ok := messages[code]; ok {
		err.Message = msg
	}
	return &err
}

func ThrowServerError(code int) error {
	err := APIError{
		Code:    500,
		Message: "SERVER ERROR",
	}
	if msg, ok := messages[code]; ok {
		err.Message = msg
	}
	return &err
}

func GetErrorMessage(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "Something Went wrong"
}

func NewError(format string, a ...interface{}) error {
	msg := fmt.Sprintf(format, a...)
	return errors.New(msg)
}
