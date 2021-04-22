package response

import "net/http"

type Info struct {
	Status       string
	ResponseCode string
	Message      string
}

type Errors struct {
	UserErrors     []string
	EngineerErrors []string
}

type Response struct {
	Status         string      `json:"status"`
	ResponseCode   string      `json:"response_code"`
	Message        string      `json:"message"`
	UserErrors     []string    `json:"user_errors"`
	EngineerErrors []string    `json:"developer_errors"`
	Signature      string      `json:"signature"`
	Result         interface{} `json:"result"`
}

func NewResponse(
	info Info,
	errors Errors,
	signature string,
	result interface{}) Response {
	return Response{
		Status:         info.Status,
		ResponseCode:   info.ResponseCode,
		Message:        info.Message,
		UserErrors:     errors.UserErrors,
		EngineerErrors: errors.EngineerErrors,
		Signature:      signature,
		Result:         result,
	}
}

func BuildErrors(userError error, engineerError error) Errors {
	var userErrors, engineerErrors []string
	userErrors = append(userErrors, userError.Error())
	engineerErrors = append(engineerErrors, engineerError.Error())
	return Errors{UserErrors: userErrors, EngineerErrors: engineerErrors}
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

var Mapping = map[string]Info{
	"ok":         {StatusSuccess, "20000000", http.StatusText(http.StatusOK)},
	"created":    {StatusSuccess, "20100000", http.StatusText(http.StatusCreated)},
	"badRequest": {StatusError, "40000000", http.StatusText(http.StatusBadRequest)},
}
