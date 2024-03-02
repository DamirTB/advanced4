package contact

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/repository"
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
	query := `INSERT INTO contacts (name, surname, patronymic, phone) VALUES ($1, $2, $3, $4)`
	res, err := r.db.Exec(query, contact.Name, contact.Surname, contact.Patronymic, contact.Phone)
	if err != nil {
		return 0, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (r *ContactRepository) GetByID(id int) (contact.Contact, error) {
	query := `SELECT * FROM contacts WHERE id = $1`
	var c contact.Contact
	if err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Surname, &c.Patronymic, &c.Phone); err != nil {
		return contact.Contact{}, err
	}

	return c, nil
}

func (r *ContactRepository) Update(contact contact.Contact) error {
	query := ` UPDATE contacts SET name = $1, surname = $2, patronymic = $3, phone = $4 WHRERE id = $5`
	_, err := r.db.Exec(query, contact.Name, contact.Surname, contact.Patronymic, contact.Phone, contact.ID)
	return err
}

func (r *ContactRepository) Delete(id int) error {
	query := `DELETE FROM contacts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
