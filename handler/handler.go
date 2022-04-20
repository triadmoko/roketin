package handler

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"path/filepath"
	"roketin/formatter"
	"roketin/helper"
	"roketin/input"
	"roketin/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}
func (h *handler) Create(c *gin.Context) {
	file, err := c.FormFile("filename")
	fileExtenstion := filepath.Ext(file.Filename)
	switch fileExtenstion {
	case ".mp4":
	case ".mkv":
	default:
		data := gin.H{"Upload file": false}
		response := helper.ResponApi("Failed Upload file Extenstion", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	//  string to lowecase
	s := fmt.Sprintf("%X", b)
	lowerCase := strings.ToLower(s)
	fileName := lowerCase + fileExtenstion
	path := "video/" + fileName

	input := input.Film{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Duration:    c.PostForm("duration"),
		Artist:      c.PostForm("artist"),
		Genre:       c.PostForm("genre"),
		Filename:    path,
	}
	fmt.Println(input)
	// err := c.ShouldBindJSON(&input)
	if err != nil {
		// error to object
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}

		response := helper.ResponApi("Insert Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// save file
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ResponApi("Insert Failed", 201, "error", errorMessage)
		c.JSON(201, response)
		return
	}
	film, err := h.service.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ResponApi("Insert Failed", 201, "error", errorMessage)
		c.JSON(201, response)
		return
	}
	formatJSON := formatter.FormatterFilm(film)
	response := helper.ResponApi("Insert Success", http.StatusOK, "success", formatJSON)

	c.JSON(http.StatusOK, response)
}
