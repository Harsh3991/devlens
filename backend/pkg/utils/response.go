package utils

import (
 "encoding/json"
 "net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
 Error string `json:"error"`
 Message string `json:"message"`
 Code int `json:"code"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
 Success bool `json:"success"`
 Data interface{} `json:"data"`
 Message string `json:"message,omitempty"`
}

// RespondWithError sends an error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
 RespondWithJSON(w, code, ErrorResponse{
 Error: http.StatusText(code),
 Message: message,
 Code: code,
 })
}

// RespondWithJSON sends a JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
 response, err := json.Marshal(payload)
 if err != nil {
 w.WriteHeader(http.StatusInternalServerError)
 w.Write([]byte(`{"error":"Internal Server Error"}`))
 return
 }

 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(code)
 w.Write(response)
}

// RespondWithSuccess sends a success response
func RespondWithSuccess(w http.ResponseWriter, data interface{}, message string) {
 RespondWithJSON(w, http.StatusOK, SuccessResponse{
 Success: true,
 Data: data,
 Message: message,
 })
}