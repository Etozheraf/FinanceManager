package Users

import usrs "TheGreatestFinanceManagerEver/Application/Models/Users"

type UserService interface {
	AddUser(telegramId int) usrs.User
	GetUser(login string, password string) usrs.User
	AddUserById(id int)
	DeleteUser(id int)
	UpdateUser(id int, login string, telegramId int, password string)
}
