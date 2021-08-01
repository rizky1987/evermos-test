package helper

import (
	"net/http"
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ServerResponse struct {
	Code int
	Type string
}

var (
	SuccessServerResponse                      ServerResponse = ServerResponse{200, "success"}
	BadRequestErrorServerResponse              ServerResponse = ServerResponse{400, "bad_request"}
	UnauthorizedErrorServerResponse            ServerResponse = ServerResponse{401, "unauthorized"}
	DatabaseErrorServerResponse                ServerResponse = ServerResponse{402, "database_error"}
	ForbiddenErrorServerResponse               ServerResponse = ServerResponse{403, "forbidden"}
	NotFoundServerResponse                     ServerResponse = ServerResponse{404, "not_found"}
	RequestTimeOutServerResponse               ServerResponse = ServerResponse{408, "request_time_out"}
	InternalServerErrorServerResponse          ServerResponse = ServerResponse{500, "internal_server_error"}
	NotImplementedServerResponse               ServerResponse = ServerResponse{501, "not_implemented"}
	ServiceTemporarilyOverloadedServerResponse ServerResponse = ServerResponse{502, "service_temporarily_overloaded"}
	ServiceUnavailableServerResponse           ServerResponse = ServerResponse{503, "service_unavailable"}
)

// ResponseHelper ...
type ResponseHelper struct {
	C           echo.Context
	Status      string
	ErrMessages []string
	Result      interface{}
	Code        int // not the http code
	CodeType    string
}

// HTTPHelper ...
type HTTPHelper struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func (u *HTTPHelper) getTypeData(i interface{}) string {
	v := reflect.ValueOf(i)
	v = reflect.Indirect(v)

	return v.Type().String()
}

// GetStatusCode ...
func (u *HTTPHelper) GetStatusCode(err error) int {
	statusCode := http.StatusOK
	if err != nil {
		switch u.getTypeData(err) {
		case "models.ErrorUnauthorized":
			statusCode = http.StatusUnauthorized
		case "models.ErrorNotFound":
			statusCode = http.StatusNotFound
		case "models.ErrorConflict":
			statusCode = http.StatusConflict
		case "models.ErrorInternalServer":
			statusCode = http.StatusInternalServerError
		default:
			statusCode = http.StatusInternalServerError
		}
	}

	return statusCode
}

// SetResponse ...
// Set response data.
func (u *HTTPHelper) SetResponse(c echo.Context, errMessages []string, result interface{}, serverResponse ServerResponse) ResponseHelper {
	return ResponseHelper{c, serverResponse.Type, errMessages, result, serverResponse.Code, serverResponse.Type}
}

// SendError ...
// Send error response to consumers.
func (u *HTTPHelper) SendError(c echo.Context, errMessages []string) error {
	res := u.SetResponse(c, errMessages, nil,  ServiceUnavailableServerResponse)

	return u.SendResponse(res)
}

// SendBadRequest ...
// Send bad request response to consumers.
func (u *HTTPHelper) SendBadRequest(c echo.Context, errMessages []string) error {

	res := u.SetResponse(c, errMessages, nil, BadRequestErrorServerResponse)

	return u.SendResponse(res)
}

// SendNotFoundRequest ...
// Send bad request response to consumers.
func (u *HTTPHelper) SendNotFoundRequest(c echo.Context, errMessages []string) error {

	res := u.SetResponse(c, errMessages, nil, NotFoundServerResponse)

	return u.SendResponse(res)
}

// SendValidationError ...
// Send validation error response to consumers.
func (u *HTTPHelper) SendValidationError(c echo.Context, validationErrors validator.ValidationErrors) error {
	errorResponse := []string{}
	errorTranslation := validationErrors.Translate(u.Translator)
	for _, err := range validationErrors {
		errorResponse = append(errorResponse, errorTranslation[err.Namespace()])
	}

	res := u.SetResponse(c, errorResponse, nil, BadRequestErrorServerResponse)

	return u.SendResponse(res)
}

// SendDatabaseError ...
// Send database error response to consumers.
func (u *HTTPHelper) SendDatabaseError(c echo.Context, message []string, errMessages []string) error {
	res := u.SetResponse(c, errMessages, nil, DatabaseErrorServerResponse)

	return u.SendResponse(res)
}

// SendUnauthorizedError ...
// Send unauthorized response to consumers.
func (u *HTTPHelper) SendUnauthorizedError(c echo.Context, errMessages []string) error {
	res := u.SetResponse(c,  errMessages, nil, UnauthorizedErrorServerResponse)

	return u.SendResponse(res)
}


// SendSuccess ...
// Send success response to consumers.
func (u *HTTPHelper) SendSuccess(c echo.Context, dataResult interface{}) error {
	res := u.SetResponse(c, []string{},dataResult, SuccessServerResponse)

	return u.SendResponse(res)
}

// Error Middleware
//Error Middleware
func SendErrorMiddleware(c echo.Context, message string, serverResponse ServerResponse) error {
	return c.JSON(serverResponse.Code, map[string]interface{}{
		"code":      serverResponse.Code,
		"code_type": serverResponse.Type,
		"message":   []string{message},
	})
}

// SendResponse ...
// Send response
func (u *HTTPHelper) SendResponse(res ResponseHelper) error {
	if len(res.ErrMessages) < 1 {
		res.ErrMessages = []string{"success"}
	}

	if res.Code != 200 {

		return res.C.JSON(res.Code, map[string]interface{}{
			"code"				: res.Code,
			"code_type"			: res.CodeType,
			"error_messages"	: res.ErrMessages,
		})
	}

	return res.C.JSON(res.Code, map[string]interface{}{
		"code"				: res.Code,
		"code_type"			: res.CodeType,
		"error_messages"	: []string{},
		"result"			: res.Result,
	})


}

func (u *HTTPHelper) EmptyJsonMap() map[string]interface{} {
	return nil //make(map[string]interface{})
}

