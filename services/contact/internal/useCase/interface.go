package usecase

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/domain/group"
)

type InterfaceContactUsecase interface {
	CreateContact(contact.Contact) (int, error) 
	GetContact(int) (contact.Contact, error)
	UpdateContact(contact.Contact) error
	DeleteContact(int) error 
}

type InterfaceGroupUsecase interface {
	CreateGroup(group.Group) (int, error) 
	GetGroup() (group.Group, error)
	InsertContact(contact.Contact, int) error 
}