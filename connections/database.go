package connections

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/yugarinn/pigeon-api/utils"
)

const projectDirName = "pigeon-api"

func Database() *gorm.DB {
	utils.LoadEnvFile(os.Getenv("PIGEON_ENV"))

	databaseURI := os.Getenv("DB_URI")
	database, err := gorm.Open(mysql.Open(databaseURI), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	return database
}
