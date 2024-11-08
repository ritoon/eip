package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/model"
)

// CreateAddress godoc
// @Summary create address
// @Description create address
// @Tags address
// @Accept  json
// @Produce  json
// @Param address body model.Address true "Address"
// @Success 201 {object} model.Address
// @Router /addresses [post]
func (h *Handler) CreateAddress(ctx *gin.Context) {
	// create a new address and bind it to the context body
	var u model.Address
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create the address in the database
	err = h.db.CreateAddress(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// return the address created
	ctx.JSON(http.StatusCreated, u)
}

// GetAddress godoc
// @Summary get address
// @Description get address
// @Tags address
// @Accept  json
// @Produce  json
// @Param uuid path string true "Address UUID"
// @Success 200 {object} model.Address
// @Router /addresses/{uuid} [get]
func (h *Handler) GetAddress(ctx *gin.Context) {
	// get the address by uuid parameter from the url
	uuid := ctx.Param("uuid")
	// get the address from the database
	u, err := h.db.GetAddress(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// return the address
	ctx.JSON(http.StatusOK, u)
}

// UpdateAddress
// @Summary update Address
// @Schemes http
// @Tags Address
// @Accept json
// @Param uuid path string true "uuid of the Address"
// @Param Authorization header string true "
// @Param Address body model.Address true "Address"
// @Produce json
// @Success 201 {object} model.Address
// @Router /addresss/:uuid [patch]
func (h *Handler) UpdateAddress(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	// create a new map and bind it to the context body
	payload := make(map[string]interface{})
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the Address in the database with the given uuid and payload
	err = h.db.UpdateAddress(uuid, payload)
	if err != nil {
		RespErr(ctx, err)
		return
	}

	u, err := h.db.GetAddress(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}

	// return the Address created
	ctx.JSON(http.StatusCreated, u)
}

// GetAddresses godoc
// @Summary get addresses
// @Description get addresses
// @Tags address
// @Accept  json
// @Produce  json
// @Success 202 {array} model.Address
// @Router /addresses/{uuid} [delete]
func (h *Handler) DeleteAddress(ctx *gin.Context) {
	// get the address by uuid parameter from the url
	uuid := ctx.Param("uuid")
	// delete the address from the database
	err := h.db.DeleteAddress(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// return nil in the response body
	ctx.JSON(http.StatusAccepted, nil)
}
