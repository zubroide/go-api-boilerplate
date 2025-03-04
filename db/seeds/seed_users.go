package seeds

import (
	"github.com/zubroide/go-api-boilerplate/model/entity"
	"gorm.io/gorm"
)

type SeedUsers struct {
	SeederBase
}

func (s *SeedUsers) Seed(db *gorm.DB) error {
	values := []*entity.User{
		{Name: "Test user"},
	}

	for _, value := range values {
		db.Where(entity.User{Name: value.Name}).
			Assign(entity.User{
				Name: value.Name,
			}).
			FirstOrCreate(value)
	}

	return nil
}
