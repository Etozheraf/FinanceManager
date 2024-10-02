package Repositories

import (
	trnscts "TheGreatestFinanceManagerEver/Application/Models/Transactions"
)

type TransactionsRepository interface {
	AddTransaction(sum int, description string, category string, userId int) trnscts.Transaction
	DeleteTransaction(transactionId int)
	GetByUser(userId int) []trnscts.Transaction
	GetByUserLastMonth(userId int) []trnscts.Transaction
	UpdateTransaction(transactionId int, description string, category string, userId int) trnscts.Transaction
}
