package Repositories

import (
	usrs "TheGreatestFinanceManagerEver/Application/Models/Users"
)

type UserRepository interface {
	GetUser(telegramId int) usrs.User
	AddUser(login string, password string) usrs.User
	AddUserById(id int) usrs.User
	DeleteUser(id int)
	UpdateUser(id int, login string, telegramId int, password string) usrs.User
}
