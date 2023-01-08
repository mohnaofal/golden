package response

import (
	"encoding/json"
	"net/http"
	"strconv"

	randomid "github.com/mohnaofal/golden/helper/generator/random_id"
)

const (
	// JSONContentType default JSON mime type
	JSONContentType = "application/json"
	// JSONCharset default JSON charset
	JSONCharset = "utf-8"
)

type Response struct {
	Error   bool        `json:"error"`
	ReffID  string      `json:"reff_id"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Write write raw data to response writer
func Write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", JSONContentType)
	w.WriteHeader(status)
	if data == nil {
		return
	}

	content, _ := json.Marshal(data)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	_, _ = w.Write(content)
}

// Error write error content with id and multilang message
func Error(w http.ResponseWriter, status int, message string) {
	content := Response{Error: true, Message: message, ReffID: randomid.GenerateID()}
	Write(w, status, content)
}

// Success func response
func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	content := Response{Error: false, Message: message, ReffID: randomid.GenerateID(), Data: data}
	Write(w, status, content)
}
