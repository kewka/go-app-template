package httphandler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/kewka/go-app-template/pkg/app"
)

// ErrorResponse ...
type ErrorResponse struct {
	Message string `json:"message"`
}

// Error ...
func Error(w http.ResponseWriter, statusCode int, err error) error {
	return JSON(w, statusCode, ErrorResponse{Message: err.Error()})
}

// JSON ...
func JSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}
	return nil
}

// BindJSON ...
func BindJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Pagination ...
func Pagination(r *http.Request) app.Pagination {
	limit, _ := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
	offset, _ := strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	return app.Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
