package user

import (
	"github.com/google/uuid"
	"go-echo-api/package/entity"
)

type IRepo struct {
	m map[string]*entity.User
}

func NewInMemRepository() *IRepo {
	var m = map[string]*entity.User{}
	return &IRepo{
		m: m,
	}
}

func (r *IRepo) Store(u *entity.User) (uuid.UUID, error) {
	r.m[u.ID.String()] = u
	return u.ID, nil
}

func (r *IRepo) FindAll() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}
