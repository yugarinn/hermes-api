package factories

import (
	"github.com/yugarinn/hermes-api/app/users/inputs"
	"github.com/yugarinn/hermes-api/app/users/models"
	"github.com/jaswdr/faker"
)

func CreateUser(properties inputs.CreateUserInput) users.User {
	faker := faker.New()
    user := users.User{Email: properties.Email, Name: properties.Name, LastName: properties.LastName, PhoneNumber: properties.PhoneNumber, CountryCode: properties.CountryCode}

	if len(user.Email) == 0 {
		user.Email = faker.Person().Contact().Email
	}

	if len(user.PhoneNumber) == 0 {
		user.PhoneNumber = faker.Person().Contact().Phone
	}

	if len(user.CountryCode) == 0 {
		user.CountryCode = "ES"
	}

	if len(user.Name) == 0 {
		user.Name = faker.Person().Name()
	}

	if len(user.LastName) == 0 {
		user.Name = faker.Person().LastName()
	}

	database.Create(&user)

	return user
}

func CreateUsersList(number int, properties inputs.CreateUserInput) []users.User {
	var users []users.User

	for i := 0; i < number; i++ {
		users = append(users, CreateUser(properties))
	}

	return users
}
