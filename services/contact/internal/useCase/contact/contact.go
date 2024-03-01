package contact

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/repository"
	usecase "arch/services/contact/internal/useCase"
)

type ContactUsecase struct {
	contactRepo repository.InterfaceContactRepository
}

func NewContactUsecase(repo repository.InterfaceContactRepository) usecase.InterfaceContactUsecase {
	return &ContactUsecase{
		contactRepo: repo,
	}
}

func (cs *ContactUsecase) CreateContact(contact contact.Contact) (int, error) {
	return 0, nil
}

func (cs *ContactUsecase) GetContact(ID int) (contact.Contact, error) {
	return contact.Contact{}, nil
}

func (cs *ContactUsecase) UpdateContact(contact contact.Contact) error {
	return nil
}

func (cs *ContactUsecase) DeleteContact(int) error {
	return nil
}
