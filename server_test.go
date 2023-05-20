package go_http_hello_world

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHelloServer(t *testing.T) {
	var path_input = "/test/123"
	var path_output = "Path: '/test/123'\n"

	req := httptest.NewRequest(http.MethodGet, "/"+path_input, nil)
	w := httptest.NewRecorder()

	HelloServer(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != path_output {
		t.Errorf("expected '%s' got '%v'", path_output, string(data))
	}
}
