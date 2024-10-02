package Application

import (
	repos "TheGreatestFinanceManagerEver/Application/Abstractions/Repositories"
	usrs "TheGreatestFinanceManagerEver/Application/Models/Users"
)

type UserServiceImpl struct {
	user                  *usrs.User
	transactionRepository *repos.TransactionsRepository
}
