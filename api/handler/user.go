package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/geocoding"
	"github.com/ritoon/eip/api/model"
	"github.com/ritoon/eip/api/util"
)

var dbConn = db.New()

// @BasePath /api/v1

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
	var payload model.UserLogin
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = payload.ValidateLogin()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := dbConn.GetUserByEmail(payload.Email)
	if err != nil {
		RespErr(ctx, err)
		return
	}

	if u.Pass == nil || payload.Pass == nil || *u.Pass != *payload.Pass {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	jwtValue, err := util.NewJWT(u.UUID, u.Email)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if u.Address.Zip != "" && u.Address.City != "" && u.Address.Street != "" {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		lat, lng, err := geocoding.New(ctxWithTimeout, "")
		if err != nil {
			if !geocoding.ErrIsTimeout(err) {
				RespErrWithCode(ctx, http.StatusServiceUnavailable, err)
				return
			}
		}
		u.Address.Lat = lat
		u.Address.Lng = lng
	}

	err = dbConn.CreateUser(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

// GetUser godoc
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
	u, err := dbConn.GetUser(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// u.Pass = nil
	ctx.JSON(http.StatusOK, u)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := dbConn.DeleteUser(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}
