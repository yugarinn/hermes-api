package tests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yugarinn/pigeon-api/app/users/models"
	"github.com/yugarinn/pigeon-api/app/messages/models"
	"github.com/yugarinn/pigeon-api/connections"
	"github.com/yugarinn/pigeon-api/http"
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

func AuthenticateAs(userId uint64, request *http.Request) {
	var user users.User
	database.Where("id", userId).First(&user)
	accessToken := "access_token"
	// TODO
	// accessToken, _ := authManager.GenerateAccessTokenForUser(user)

	request.Header.Set("Authorization", accessToken)
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

func databaseSetup() {
	database.AutoMigrate(&messages.Message{})
}

func databaseTeardown() {
	database.Migrator().DropTable(&messages.Message{})
}
