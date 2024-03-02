package usecase

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/domain/group"
)

type IContactUsecase interface {
	Create(contact.Contact) (int, error)  
	GetByID(int) (contact.Contact, error) 
	Delete(int) error

	CreateGroup(group.Group) (int, error)
	DeleteGroup(int) error                
	UpdateGroup(group.Group) error
	GetGroupByID(int) (group.Group, error) 
	InsertContactToGroup(int, int) error   
	DeleteContactFromGroup(int, int) error 
}

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