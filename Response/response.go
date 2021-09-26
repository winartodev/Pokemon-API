package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BodyError struct {
	Code   int    `json:"http_code"`
	Status string `json:"http_status"`
	Error  string `json:"error"`
}

type BodySuccess struct {
	Code    int         `json:"http_code"`
	Status  string      `json:"http_status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RespondWithJSON(w http.ResponseWriter, result interface{}, code int) {
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, string(response))
}

// NewResponseError will return if response error
func NewResponseError(result *BodyError) *BodyError {
	return &BodyError{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}

// NewResponseSuccess will return if response success
func NewResponseSuccess(data interface{}) *BodySuccess {
	return &BodySuccess{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success",
		Data:    data,
	}
}

// NewPokemonCreateSuccess will return if pokemon success create
func NewPokemonCreateSuccess(id int) *BodySuccess {
	return &BodySuccess{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: "Success",
		Data:    map[string]string{"status": fmt.Sprintf("id pokemon %d create", id)},
	}
}
