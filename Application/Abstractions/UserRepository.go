package Abstractions

type UserRepository interface {
	GetUser(telegram_id int)
	AddUser(login string, password string)
	AddUserById(id int)
	DeleteUser(login string)
}
