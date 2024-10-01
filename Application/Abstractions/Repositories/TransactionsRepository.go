package Repositories

type TransactionsRepository interface {
	AddTransaction(sum int, description string, category string, userId int)
	DeleteTransaction(transactionId int)
	GetByUser(userId int)
	GetByUserLastMonth(userId int)
}
