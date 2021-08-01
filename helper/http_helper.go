package helper

import (
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

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
	C        echo.Context
	Status   string
	Message  []string
	Data     interface{}
	Code     int // not the http code
	CodeType string
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
func (u *HTTPHelper) SetResponse(c echo.Context, status string, message []string, data interface{}, serverResponse ServerResponse) ResponseHelper {
	return ResponseHelper{c, status, message, data, serverResponse.Code, serverResponse.Type}
}

// SendError ...
// Send error response to consumers.
func (u *HTTPHelper) SendError(c echo.Context, message []string, data interface{}) error {
	res := u.SetResponse(c, `error`, message, data, BadRequestErrorServerResponse)

	return u.SendResponse(res)
}

// SendBadRequest ...
// Send bad request response to consumers.
func (u *HTTPHelper) SendBadRequest(c echo.Context, message string, data interface{}) error {

	res := u.SetResponse(c, `error`, []string{message}, data, BadRequestErrorServerResponse)

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

	res := u.SetResponse(c, `error`, errorResponse, nil, BadRequestErrorServerResponse)

	return u.SendResponse(res)
}

// SendDatabaseError ...
// Send database error response to consumers.
func (u *HTTPHelper) SendDatabaseError(c echo.Context, message []string, data interface{}) error {
	return u.SendError(c, message, data)
}

// SendUnauthorizedError ...
// Send unauthorized response to consumers.
func (u *HTTPHelper) SendUnauthorizedError(c echo.Context, message []string, data interface{}) error {
	return u.SendError(c, message, data)
}

// SendNotFoundError ...
// Send not found response to consumers.
func (u *HTTPHelper) SendBadFoundError(c echo.Context, message []string, data interface{}) error {
	return u.SendError(c, message, data)
}

// SendSuccess ...
// Send success response to consumers.
func (u *HTTPHelper) SendSuccess(c echo.Context, data interface{}) error {
	res := u.SetResponse(c, `ok`, []string{}, data, SuccessServerResponse)

	return u.SendResponse(res)
}

// SendResponse ...
// Send response
func (u *HTTPHelper) SendResponse(res ResponseHelper) error {
	if len(res.Message) == 0 {
		res.Message = append(res.Message, `success`)
	}

	var resCode int
	if res.Code != 200 {
		resCode = http.StatusBadRequest
	} else {
		resCode = http.StatusOK
	}

	return res.C.JSON(resCode, map[string]interface{}{
		"code":      res.Code,
		"code_type": res.CodeType,
		"message":   res.Message,
		"data":      res.Data,
	})
}

func (u *HTTPHelper) EmptyJsonMap() map[string]interface{} {
	return nil //make(map[string]interface{})
}

//get pagination URL
func (u *HTTPHelper) GetPagingUrl(c echo.Context, page, limit int) string {

	r := c.Request()
	currentURL := c.Scheme() + "://" + r.Host + r.URL.Path + "?page={page}&limit={limit}"

	defaultLinkReplacer := strings.NewReplacer("{page}", strconv.Itoa(page), "{limit}", strconv.Itoa(limit)).Replace(currentURL)

	return defaultLinkReplacer
}

//Set paginantion response
func (u *HTTPHelper) GeneratePaging(c echo.Context, prev, next, limit, page, totalRecord int) map[string]interface{} {

	prevURL, nextURL, firstURL, lastURL := "", "", "", ""
	paramPrevURL, paramNextURL, paramFirstURL, paramLastURL := "", "", "", ""

	totalPages := int(math.Ceil(float64(totalRecord) / float64(limit)))

	if page >= 1 {
		prev = page - 1
		if page < totalPages {
			next = page + 1
		} else {
			next = totalPages
		}
	}

	if totalPages >= page && page > 1 {
		prevURL = u.GetPagingUrl(c, prev, limit)
		paramPrevURL = "?page=" + strconv.Itoa(prev) + "&limit=" + strconv.Itoa(limit)
	}

	if totalPages > page {
		nextURL = u.GetPagingUrl(c, next, limit)
		paramNextURL = "?page=" + strconv.Itoa(next) + "&limit=" + strconv.Itoa(limit)
	}

	if totalPages >= page && page > 1 {
		firstURL = u.GetPagingUrl(c, 1, limit)
		paramFirstURL = "?page=1" + "&limit=" + strconv.Itoa(limit)
	}

	if totalPages >= page && totalPages != page {
		lastURL = u.GetPagingUrl(c, totalPages, limit)
		paramLastURL = "?page=" + strconv.Itoa(totalPages) + "&limit=" + strconv.Itoa(limit)
	}

	links := map[string]interface{}{
		"previous": prevURL,
		"next":     nextURL,
		"first":    firstURL,
		"last":     lastURL,
	}

	linkParameter := map[string]interface{}{
		"previous": paramPrevURL,
		"next":     paramNextURL,
		"first":    paramFirstURL,
		"last":     paramLastURL,
	}

	pagination := map[string]interface{}{
		"total_records":  totalRecord,
		"per_page":       limit,
		"current_page":   page,
		"total_pages":    totalPages,
		"links":          links,
		"link_parameter": linkParameter,
	}

	return pagination
}
