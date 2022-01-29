package jobs

import (
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	fname                    = "Test"
	lname                    = "User"
	nname                    = "Foo Bar"
	description              = "This is the first user ever!"
	location                 = "Their house, maybe? Wouldn't know"
	firstUser   *models.User = &models.User{
		Email:       "bibi@test.com",
		FirstName:   &fname,
		LastName:    &lname,
		NickName:    &nname,
		Description: &description,
		Location:    &location,
	}
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
