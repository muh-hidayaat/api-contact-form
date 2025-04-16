package handlers

import (
	"api-contact-form/requests"
	"api-contact-form/responses"
	"api-contact-form/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	service services.ContactService
}

func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req requests.ContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	contact, err := h.service.CreateContact(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, responses.APIResponse{
		Code:    "CREATED",
		Message: "Contact created successfully",
		Data:    responses.ContactResponseFromModel(contact),
	})
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	contacts, err := h.service.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var contactResponses []responses.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, responses.ContactResponseFromModel(&contact))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contacts retrieved successfully",
		Data:    contactResponses,
	})
}

func (h *ContactHandler) GetContact(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	contact, err := h.service.GetContactByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.APIResponse{
			Code:    "NOT_FOUND",
			Message: "Contact not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contact retrieved successfully",
		Data:    responses.ContactResponseFromModel(contact),
	})
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var req requests.ContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	contact, err := h.service.UpdateContact(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contact updated successfully",
		Data:    responses.ContactResponseFromModel(contact),
	})
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	err = h.service.DeleteContact(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contact deleted successfully",
		Data:    nil,
	})
}
