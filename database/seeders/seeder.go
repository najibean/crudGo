package seeders

import (
	"github.com/najibean/crudGo/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	seeder interface{}
}

// return nya adalah array dari seeder
func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{seeder: fakers.UserFaker(db)},
		{seeder: fakers.ProductFaker(db)},
	}
}

// melooping seeder yang dibuat dari faker untuk dimasukkan kedalam database
func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
