package Abstractions

type UserRepository interface {
	GetUser(telegramId int)
	AddUser(login string, password string)
	AddUserById(id int)
	DeleteUser(login string)
}
