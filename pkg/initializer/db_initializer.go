package initializer

// TODO: move initializers to config?
// export from config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mccuskero/go-user-management-sandbox/pkg/config"
	"github.com/mccuskero/go-user-management-sandbox/pkg/models"
)

type Initializer struct {
	DB *gorm.DB
}

func NewInitializer() *Initializer {
	return &Initializer{
		DB: nil,
	}
}

func (in *Initializer) ConnectDB(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)

	in.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")
}

func (in *Initializer) AutoMigrate() {
	// automigrate the tables (maybe we disable this in the future)
	// Automigrate the User model
	// AutoMigrate() automatically migrates our schema, to keep our schema upto date.
	err := in.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to AutoMigrate schema")
	}
}
