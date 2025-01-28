package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ParseBody(r *http.Request, input any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, input); err != nil {
		return err
	}
	return nil
}

func ParseQuery(query string) map[string]string {
	result := make(map[string]string)

	optionsArray := strings.Split(query, "&")
	for _, option := range optionsArray {
		variable, value, _ := strings.Cut(option, "=")

		result[variable] = value
	}
	return result
}

func ParseQueryID(r *http.Request, key string) (uint64, error) {
	var err error

	request := ParseQuery(r.URL.RawQuery)
	requestID, ok := request[key]
	if !ok {
		return 0, fmt.Errorf("requestID | error with %v : %v", key, fmt.Errorf("no %s query parameter found", key))
	}

	id, err := strconv.ParseUint(requestID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("ParseUint | error with %v: %v", key, err)
	}

	return id, nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{Error: err.Error()}

	jsonResponse, _ := json.Marshal(response)
	_, _ = w.Write(jsonResponse)
}

func SendJSON(w http.ResponseWriter, status int, output any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	data, err := json.Marshal(output)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err)
		return
	}
	_, _ = w.Write(data)
}

func SendRAWJSON(w http.ResponseWriter, status int, output []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	_, _ = w.Write(output)
}

func SendCode(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
