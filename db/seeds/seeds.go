package seeds

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Seeds manager seeders
type Seeds struct {
	db    *gorm.DB
	items []Seeder
}

// Seeder seeder interface
type Seeder interface {
	GetName() string
	Seed(db *gorm.DB) error
}

type SeederBase struct {
	Name string
}

func (s *SeederBase) GetName() string {
	return s.Name
}

// NewSeeds factory new seeders manager
func NewSeeds(db *gorm.DB) *Seeds {
	seeds := new(Seeds)
	seeds.db = db
	return seeds
}

// AppendSeeder register seeder
func (s *Seeds) AppendSeeder(seeder Seeder) error {
	if nil != s.FindSeeder(seeder.GetName()) {
		err := fmt.Sprintf("Seeder with name (%s) already exists", seeder.GetName())
		return errors.New(err)
	}
	s.items = append(s.items, seeder)
	return nil
}

func (s *Seeds) FindSeeder(name string) Seeder {
	for _, seeder := range s.items {
		if seeder.GetName() == name {
			return seeder
		}
	}
	return nil
}

// RunSeeder run target seeder
func (s *Seeds) RunSeeder(seeder Seeder) error {
	fmt.Printf("Seeder (%s) start\n", seeder.GetName())
	// run seeder
	if err := seeder.Seed(s.db); nil != err {
		return err
	}
	fmt.Printf("Seeder (%s) complete\n", seeder.GetName())

	return nil
}

// RunSeederByName run target seeder by name
func (s *Seeds) RunSeederByName(name string) error {
	seeder := s.FindSeeder(name)
	if nil == seeder {
		err := fmt.Sprintf("Seeder with name (%s) does not exists", name)
		return errors.New(err)
	}
	return s.RunSeeder(seeder)
}

// RunSeeds run all seeders
func (s *Seeds) RunSeeds() error {
	for _, seeder := range s.items {
		if err := s.RunSeeder(seeder); nil != err {
			return err
		}
	}
	return nil
}
