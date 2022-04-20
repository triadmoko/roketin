package handler

import (
	"net/http"
	"roketin/formatter"
	"roketin/helper"
	"roketin/input"
	"roketin/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}
func (h *handler) Create(c *gin.Context) {
	var input input.Film

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// error to object
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}

		response := helper.ResponApi("Insert Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
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
