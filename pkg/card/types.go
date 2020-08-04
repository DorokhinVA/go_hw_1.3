package card

import "time"

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	Amount int64
	Date   time.Time
	MCC    string
	Status string
	Type   TransactionType
}

type TransactionType int8

const (
	DEPOSIT  TransactionType = 1
	WITHDRAW TransactionType = 2
)
