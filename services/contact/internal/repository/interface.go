package repository

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/domain/group"
)

type InterfaceContactRepository interface {
	Insert(contact.Contact) (int, error)  
	GetByID(int) (contact.Contact, error) 
	Update(contact.Contact) error
	Delete(int) error 
}

type InterfaceGroupRepository interface {
	Insert(group.Group) (int, error)          
	GetByID(int) (group.Group, error)         
	InsertContact(contact.Contact, int) error 
}