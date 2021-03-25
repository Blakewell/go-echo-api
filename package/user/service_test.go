package user

import (
	"github.com/stretchr/testify/assert"
	"go-echo-api/package/entity"
	"testing"
)

func TestFindAll(t *testing.T) {

	repo := NewInMemRepository()
	service := NewService(repo)

	u1 := &entity.User{Name: "Blake"}
	u2 := &entity.User{Name: "Kacee"}

	_, _ = service.Store(u1)
	_, _ = service.Store(u2)

	t.Run("find all", func(t *testing.T) {
		all, err := service.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})
}
