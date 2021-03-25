package user

import (
	"github.com/google/uuid"
	"go-echo-api/package/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) FindAll() ([]*entity.User, error) {
	return s.repo.FindAll()
}

func (s *Service) Store(u *entity.User) (uuid.UUID, error) {
	u.ID, _ = uuid.NewUUID()

	return s.repo.Store(u)
}
