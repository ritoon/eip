package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PageLogin(ctx *gin.Context) {
	var BaseURL = "http://localhost:8888"
	var m = map[string]interface{}{
		"BaseURL": BaseURL,
	}
	ctx.HTML(http.StatusOK, "login.html", m)
}
