package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//Postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/LandazuriPaul/good-resolutions/api/internal/orm/migration"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/logger"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
)

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and connection string
func Factory(config *utils.ServerConfig) (*ORM, error) {
	logger.Info(fmt.Sprintf("Connecting to database using following config: dialect=%s host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Database.Dialect, config.Database.Host, config.Database.Port, config.Database.Username, config.Database.DatabaseName, config.Database.Password, config.Database.SSLMode))
	db, err := gorm.Open(config.Database.Dialect, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.DatabaseName, config.Database.Password, config.Database.SSLMode))
	if err != nil {
		logger.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(config.Database.LogMode)
	// Automigrate tables
	if config.Database.AutoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	logger.Info("[ORM] Database connection initialized.")
	return orm, err
}
