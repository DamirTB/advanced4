package contact

import (
	"arch/services/contact/internal/repository"
	"arch/services/contact/internal/domain/contact"
	"database/sql"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) repository.InterfaceContactRepository {
	return &ContactRepository{
		db: db,
	}
}

func (r *ContactRepository) Insert(contact contact.Contact) (int, error) {
	return 0, nil
}

func (r *ContactRepository) GetByID(id int) (contact.Contact, error) {
	return contact.Contact{}, nil
}

func (r *ContactRepository) Update(contact contact.Contact) error {
	return nil
}

func (r *ContactRepository) Delete(id int) error {
	return nil
}