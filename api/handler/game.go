package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/model"
)

// CreateGame create game
// @Summary create game
// @Description create game
// @Tags game
// @Accept  json
// @Produce  json
// @Param game body model.Game true "Game"
// @Success 201 {object} model.Game
// @Router /games [post]
func (h *Handler) CreateGame(ctx *gin.Context) {
	// create a new game and bind it to the context body
	var u model.Game
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create the game in the database
	err = h.db.CreateGame(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// return the game created
	ctx.JSON(http.StatusCreated, u)
}

// GetGame get game
// @Summary get game
// @Description get game
// @Tags game
// @Accept  json
// @Produce  json
// @Param uuid path string true "Game UUID"
// @Success 200 {object} model.Game
// @Router /games/{uuid} [get]
func (h *Handler) GetGame(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := h.db.GetGame(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	// u.Pass = nil
	ctx.JSON(http.StatusOK, u)
}

// UpdateGame update game
// @Summary update game
// @Description update game
// @Tags game
// @Accept  json
// @Produce  json
// @Param uuid path string true "Game UUID"
// @Param game body model.Game true "Game"
// @Success 202 {object} model.Game
// @Router /games/{uuid} [delete]
func (h *Handler) DeleteGame(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := h.db.DeleteGame(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

// SearchGames search games by name
// @Summary search games by name
// @Description search games by name
// @Tags game
// @Accept  json
// @Produce  json
// @Param name query string false "Game name"
// @Success 200 {array} model.Game
// @Router /games [get]
func (h *Handler) SearchGames(ctx *gin.Context) {
	name := ctx.Query("name")
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	games, err := h.db.SearchGames(ctxWithTimeout, name)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, games)
}

// AddImageToGame add image to game
// @Summary Add image to game
// @Description Add image to game
// @Tags game
// @Accept  multipart/form-data
// @Produce  json
// @Param uuid path string true "Game UUID"
// @Param file formData file true "Image file"
// @Success 202 {object} model.Game
// @Router /games/{uuid}/image [post]
func (h *Handler) AddImageToGame(ctx *gin.Context) {
	// get the game by uuid parameter from the url
	uuid := ctx.Param("uuid")

	// get the game from the database
	_, err := h.db.GetGame(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}

	// get the file from the form
	file, _ := ctx.FormFile("file")

	// Upload the file to specific dst.
	filePath := "./uploads/" + file.Filename
	ctx.SaveUploadedFile(file, filePath)

	// read mimetype
	var buff []byte = make([]byte, 512)
	fileOpen, _ := os.Open(filePath)
	defer fileOpen.Close()
	fileOpen.Read(buff)

	// detect mimetype
	mtype := mimetype.Detect(buff)
	fmt.Println("mtype :", mtype.String(), mtype.Extension())

	// check if mimetype is image
	if mtype.String() != "image/jpeg" && mtype.String() != "image/png" {
		os.Remove(filePath)
		err := errors.New("invalid image type got: " + mtype.String())
		RespErr(ctx, err)
		return
	}

	// update game with image
	game, err := h.db.UpdateImage(uuid, filePath)
	if err != nil {
		RespErr(ctx, err)
		return
	}

	// return the game
	ctx.JSON(http.StatusAccepted, game)
}
