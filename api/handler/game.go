package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gabriel-vasile/mimetype"
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
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	games, err := dbConn.SearchGames(ctxWithTimeout, name)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, games)
}

func (h *Handler) AddImageToGame(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, err := dbConn.GetGame(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	filePath := "./uploads/" + file.Filename
	ctx.SaveUploadedFile(file, filePath)

	// read mimetype
	var buff []byte = make([]byte, 512)
	fileOpen, _ := os.Open(filePath)
	defer fileOpen.Close()
	fileOpen.Read(buff)
	log.Println("buff :", string(buff))
	mtype := mimetype.Detect(buff)
	fmt.Println("mtype :", mtype.String(), mtype.Extension())
	if mtype.String() != "image/jpeg" && mtype.String() != "image/png" {
		os.Remove(filePath)
		err := errors.New("invalid image type got: " + mtype.String())
		RespErr(ctx, err)
		return
	}
	// update game with image
	game, err := dbConn.UpdateImage(uuid, filePath)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, game)
}
