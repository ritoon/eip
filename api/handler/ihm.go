package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageLogin render login page with form to login.
func (h *Handler) PageLogin(ctx *gin.Context) {
	dataLogin := map[string]interface{}{
		"Title":   "Login",
		"BaseURL": "http://localhost:8888",
	}
	ctx.HTML(http.StatusOK, "login.html", dataLogin)
}

func (h *Handler) PageGameList(ctx *gin.Context) {

	dataLogin := map[string]interface{}{
		"Title":   "Liste des jeux",
		"BaseURL": "http://localhost:8888",
	}
	ctx.HTML(http.StatusOK, "gamelist.html", dataLogin)
}
