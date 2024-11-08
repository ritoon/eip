package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/geocoding"
	"github.com/ritoon/eip/api/model"
	"github.com/ritoon/eip/api/util"
)

// LoginUser godoc
// @Summary login user
// @Schemes http
// @Description do ping
// @Tags user
// @Accept json
// @Param payload body model.UserLogin true "User login"
// @Produce json
// @Success 200 {object} util.JWTResponse
// @Router /login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	// create a new user login payload and bind it to the context body
	var payload model.UserLogin
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate the payload and return an error if it is invalid
	err = payload.ValidateLogin()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get the user by email from the database
	u, err := h.db.GetUserByEmail(payload.Email)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// return an error if the password is invalid
	if u.Pass == nil || payload.Pass == nil || *u.Pass != *payload.Pass {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	// create a new jwt token with the user uuid and email
	jwtValue, err := util.NewJWT(u.UUID, u.Email)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// return the jwt token
	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
}

// CreateUser godoc
// @Summary create user
// @Schemes http
// @Tags user
// @Accept json
// @Param user body model.User true "User"
// @Produce json
// @Success 201 {object} model.User
// @Router /users [post]
func (h *Handler) CreateUser(ctx *gin.Context) {
	// create a new user and bind it to the context body
	var u model.User
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if the user address is not empty, get the latitude and longitude from the address
	// call the geocoding function to get the latitude and longitude
	if u.Address.Zip != "" && u.Address.City != "" && u.Address.Street != "" {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		// get the latitude and longitude from the address
		lat, lng, err := geocoding.New(ctxWithTimeout, "")
		if err != nil {
			if !geocoding.ErrIsTimeout(err) {
				RespErrWithCode(ctx, http.StatusServiceUnavailable, err)
				return
			}
		}
		// set the latitude and longitude to the user address
		u.Address.Lat = lat
		u.Address.Lng = lng
	}

	// create the user in the database with the address and games
	err = h.db.CreateUser(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}

	// return the user created
	ctx.JSON(http.StatusCreated, u)
}

// GetUser is a method that returns a user by its uuid.
// @Summary get a user by its uuid
// @Schemes http
// @Tags user
// @Accept json
// @Param uuid path string true "uuid of the user"
// @Param Authorization header string true "bearer token"
// @Success 200 {object} model.User
// @Router /users/:uuid [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := h.db.GetUser(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// u.Pass = nil
	ctx.JSON(http.StatusOK, u)
}

// DeleteUser godoc
// @Summary delete a user by its uuid
// @Schemes http
// @Tags user
// @Accept json
// @Param uuid path string true "uuid of the user"
// @Param Authorization header string true "
// @Success 202
// @Router /users/:uuid [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := h.db.DeleteUser(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}
