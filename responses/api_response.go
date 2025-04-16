package responses

import (
	"api-contact-form/helpers"
	"api-contact-form/models"
)

type APIResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ContactResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Message string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ContactResponseFromModel(contact *models.Contact) ContactResponse {
	return ContactResponse{
		ID:        contact.ID,
		Name:      contact.FullName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		Message:   contact.Message,
		CreatedAt: helpers.FormatTimeHuman(contact.CreatedAt),
		UpdatedAt: helpers.FormatTimeHuman(contact.UpdatedAt),
	}
}
