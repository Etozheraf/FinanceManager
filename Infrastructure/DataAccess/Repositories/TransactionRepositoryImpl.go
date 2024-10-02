package Repositories

import (
	"TheGreatestFinanceManagerEver/Application/Abstractions/Repositories"
	trnscts "TheGreatestFinanceManagerEver/Application/Models/Transactions"
	"database/sql"
)

type TransactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepositoryImpl(db *sql.DB) Repositories.TransactionsRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (t *TransactionRepositoryImpl) AddTransaction(sum int, description string, category string, userId int) trnscts.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionRepositoryImpl) DeleteTransaction(transactionId int) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionRepositoryImpl) GetByUser(userId int) []trnscts.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionRepositoryImpl) GetByUserLastMonth(userId int) []trnscts.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionRepositoryImpl) GetByUserLastDay(userId int) []trnscts.Transaction {
	// TODO: implement me
	panic("implement me")
}
