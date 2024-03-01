package group

import (
	"arch/services/contact/internal/domain/contact"
	"arch/services/contact/internal/domain/group"
	"arch/services/contact/internal/repository"
	usecase "arch/services/contact/internal/useCase"
)

type GroupUsecase struct {
	gropRepo repository.InterfaceGroupRepository
}

func NewGroupUseCase(repo repository.InterfaceGroupRepository) usecase.InterfaceGroupUsecase {
	return &GroupUsecase{
		gropRepo: repo,
	}
}

func (gs *GroupUsecase) CreateGroup(group.Group) (int, error) {
	return 0, nil
}

func (gs *GroupUsecase) GetGroup() (group.Group, error) {
	return group.Group{}, nil
}

func (gs *GroupUsecase) InsertContact(contact contact.Contact, id int) error {
	return nil
}