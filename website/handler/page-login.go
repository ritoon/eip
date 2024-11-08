package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageLogin render login page with form to login.
func (h *Handler) PageLogin(ctx *gin.Context) {
	dataLogin := map[string]interface{}{
		"Title":      "Login",
		"BaseURL":    "http://localhost:8880",
		"APIBaseURL": "http://localhost:8888",
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.HTML(http.StatusOK, "login.html", dataLogin)
}

func (h *Handler) PageGameList(ctx *gin.Context) {

	dataLogin := map[string]interface{}{
		"Title":      "Liste des jeux",
		"BaseURL":    "http://localhost:8880",
		"APIBaseURL": "http://localhost:8888/api/v1",
	}
	ctx.HTML(http.StatusOK, "gamelist.html", dataLogin)
}
