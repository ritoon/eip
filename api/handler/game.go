package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/model"
)

func (h *Handler) CreateGame(ctx *gin.Context) {
	var u model.Game
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = dbConn.CreateGame(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

func (h *Handler) GetGame(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := dbConn.GetGame(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// u.Pass = nil
	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) DeleteGame(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := dbConn.DeleteGame(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

func (h *Handler) SearchGames(ctx *gin.Context) {
	name := ctx.Query("name")

	games, err := dbConn.SearchGames(name)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, games)
}
