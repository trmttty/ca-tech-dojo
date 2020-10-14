package middlewear

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCors(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/cors", Cors(func(w http.ResponseWriter, r *http.Request) {
	}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("OPTIONS", "/cors", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	if header := writer.Header().Get("Access-Control-Allow-Origin"); header != "*" {
		t.Errorf("Response header is %v", header)
	}
	if header := writer.Header().Get("Access-Control-Allow-Headers"); header != "Authorization, Origin, X-Requested-With, Content-Type, Accept, x-token" {
		t.Errorf("Response header is %v", header)
	}
	if header := writer.Header().Get("Access-Control-Allow-Methods"); header != "GET, POST, PUT, DELETE, OPTIONS" {
		t.Errorf("Response header is %v", header)
	}
	if header := writer.Header().Get("Content-type"); header != "application/json" {
		t.Errorf("Response header is %v", header)
	}
}
