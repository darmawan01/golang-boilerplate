package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// RespondwithJSON help to write response
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Can't parse into json with err: ", err.Error())
	}
	w.WriteHeader(code)
	w.Write(response)
}

// DataFormat help to format data response
func DataFormat(msg string, data interface{}, total, page, limit int) interface{} {
	result := map[string]interface{}{
		"msg":   msg,
		"data":  data,
		"total": total,
		"page":  page,
		"limit": limit,
	}
	return result
}

// ErrFormat help to format error response
func ErrFormat(msg string, data interface{}) interface{} {
	result := map[string]interface{}{
		"msg":  msg,
		"data": nil,
	}
	return result
}

// GetIDParam parse id from request parameter
func GetIDParam(r *http.Request) string {
	id := mux.Vars(r)["id"]
	return id
}

// GetQueryParam Help to get query param by key
func GetQueryParam(r *http.Request, key ...string) []string {
	querys := []string{}
	for _, q := range key {
		query := r.URL.Query().Get(q)
		querys = append(querys, query)
	}
	return querys
}

// IsEmpty validate is value exist
func IsEmpty(str string) bool {
	if len(str) <= 0 || str == "" {
		return true
	}
	return false
}

// IsEqual compare two string
func IsEqual(str1 string, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

// ValidateStruct help to validate struct
func ValidateStruct(model interface{}) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(model); err != nil {
		return false, err
	}
	return true, nil
}
