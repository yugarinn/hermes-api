package tests

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yugarinn/catapi.cat/app/breeds/models"
	"github.com/yugarinn/catapi.cat/connections"
	"github.com/yugarinn/catapi.cat/http"
)


var database *gorm.DB = connections.Database()

func SetupRouter() *gin.Engine {
	gin.SetMode("test")

	router := gin.Default()
	routes.Register(router)

	return router
}

func Reset() {
	databaseTeardown()
	databaseSetup()
}

func DatabaseHas(tableName string, whereClause string) bool {
	var result bool
	database.Raw(fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s;", tableName, whereClause)).Scan(&result)

	return result
}

func DatabaseHasCount(tableName string, expectedCount int, whereClause string) bool {
	var actualCount int
	database.Raw(fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s;", tableName, whereClause)).Scan(&actualCount)

	return expectedCount == actualCount
}

func DatabaseMissing(tableName string, whereClause string) bool {
	return !DatabaseHas(tableName, whereClause)
}

// TODO: DRY
func databaseSetup() {
	database.AutoMigrate(&breeds.Breed{})
}

// TODO: DRY
func databaseTeardown() {
	database.Migrator().DropTable(&breeds.Breed{})
}
