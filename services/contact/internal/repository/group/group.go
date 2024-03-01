package group

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/domain/group"
	"arch/services/contact/internal/repository"
	"database/sql"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) repository.InterfaceGroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) Insert(group group.Group) (int, error) {
	return 0, nil
}

func (r *GroupRepository) GetByID(id int) (group.Group, error) {
	return group.Group{}, nil
}

func (r *GroupRepository) InsertContact(contact contact.Contact, id int) error {
	return nil
}