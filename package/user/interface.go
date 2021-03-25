package user

import (
	"github.com/google/uuid"
	"go-echo-api/package/entity"
)

type Reader interface {
	FindAll() ([]*entity.User, error)
}

type Writer interface {
	Store(u *entity.User) (uuid.UUID, error)
}

type Repository interface {
	Reader
	Writer
}
