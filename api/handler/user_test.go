package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	h := Handler{}
	ctx.Request = httptest.NewRequest("POST", "/login", nil)
	h.LoginUser(ctx)
	if ctx.Writer.Status() != 400 {
		t.Errorf("Expected status 400, got %d", ctx.Writer.Status())
	}
	if recorder.Body.String() != "{\"error\":\"email is required\"}" {
		t.Errorf("Expected response body {\"error\":\"email is required\"}, got %s", recorder.Body.String())
	}
}
