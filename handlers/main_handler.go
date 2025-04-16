package handlers

import (
	"api-contact-form/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainHandler struct{}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

func (h *MainHandler) MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API Contact Form is running.",
	})
}
