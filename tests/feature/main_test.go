package tests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yugarinn/hermes-api/app/users/models"
	"github.com/yugarinn/hermes-api/app/messages/models"
	"github.com/yugarinn/hermes-api/connections"
	"github.com/yugarinn/hermes-api/http"
	"github.com/yugarinn/hermes-api/core"
	"github.com/yugarinn/hermes-api/tests/mocks"
)


var database *gorm.DB = connections.Database()

func SetupRouter() (*core.App, *gin.Engine) {
	gin.SetMode("test")

	app := mockApp()
	router := gin.Default()
	routes.Register(app, router)

	return app, router
}

func ResetDatabase() {
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
	database.AutoMigrate(&users.User{})
	database.AutoMigrate(&users.UserValidation{})
	database.AutoMigrate(&messages.Message{})
}

func databaseTeardown() {
	database.Migrator().DropTable(&users.User{})
	database.Migrator().DropTable(&users.UserValidation{})
	database.Migrator().DropTable(&messages.Message{})
}

func mockApp() *core.App {
    app := &core.App{
        TwilioClient: &mocks.TwilioMock{},
    }

	return app
}
