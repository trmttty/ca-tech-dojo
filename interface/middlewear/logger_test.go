package middlewear

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogger(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/logger", Logger(func(w http.ResponseWriter, r *http.Request) {
	}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/logger", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
