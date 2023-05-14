package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sudobooo/gopher_messenger/user_service/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestService_handleHello(t *testing.T) {
	s := New(config.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", http.NoBody)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
