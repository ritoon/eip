package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/model"
)

func (h *Handler) CreateAddress(ctx *gin.Context) {
	var u model.Address
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = dbConn.CreateAddress(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

func (h *Handler) GetAddress(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := dbConn.GetAddress(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// u.Pass = nil
	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) DeleteAddress(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := dbConn.DeleteAddress(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}
