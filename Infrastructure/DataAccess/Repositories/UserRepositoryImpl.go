package Repositories

import (
	usrs "TheGreatestFinanceManagerEver/Application/Models/Users"
	"database/sql"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (u UserRepositoryImpl) GetUser(telegramId int) usrs.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) AddUser(login string, password string) usrs.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) AddUserById(id int) usrs.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) DeleteUser(id int) {

}

func (u UserRepositoryImpl) UpdateUser(id int, login string, telegramId int, password string) usrs.User {
	// TODO: implement me
	panic("implement me")
}

//GetUser(telegramId int) usrs.User
//AddUser(login string, password string) usrs.User
//AddUserById(id int) usrs.User
//DeleteUser(login string)
