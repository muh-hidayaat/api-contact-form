package services

import (
	"api-contact-form/models"
	"api-contact-form/repositories"
	"api-contact-form/requests"

	"github.com/go-playground/validator/v10"
)

type ContactService interface {
	CreateContact(req *requests.ContactRequest) (*models.Contact, error)
	GetAllContacts() ([]models.Contact, error)
	GetContactByID(id uint) (*models.Contact, error)
	UpdateContact(id uint, req *requests.ContactRequest) (*models.Contact, error)
	DeleteContact(id uint) error
}

type contactService struct {
	repository repositories.ContactRepository
	validate   *validator.Validate
}

func NewContactService(repository repositories.ContactRepository) ContactService {
	return &contactService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *contactService) CreateContact(req *requests.ContactRequest) (*models.Contact, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	contact := models.Contact{
		FullName: req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Message:  req.Message,
	}

	err := s.repository.Create(&contact)
	return &contact, err
}

func (s *contactService) GetAllContacts() ([]models.Contact, error) {
	return s.repository.FindAll()
}

func (s *contactService) GetContactByID(id uint) (*models.Contact, error) {
	return s.repository.FindByID(id)
}

func (s *contactService) UpdateContact(id uint, req *requests.ContactRequest) (*models.Contact, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	contact, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	contact.FullName = req.Name
	contact.Email = req.Email
	contact.Phone = req.Phone
	contact.Message = req.Message

	err = s.repository.Update(contact)
	return contact, err
}

func (s *contactService) DeleteContact(id uint) error {
	contact, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(contact)
}
