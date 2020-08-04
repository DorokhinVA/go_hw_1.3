package main

import (
	"fmt"
	"github.com/DorokhinVA/go_hw_1.3/pkg/card"
	"time"
)

func main() {
	firstDate, _ := time.Parse(time.RFC3339, "2020-08-03T14:08:00Z")
	secondDate, _ := time.Parse(time.RFC3339, "2020-08-03T15:08:00Z")
	thirdDate, _ := time.Parse(time.RFC3339, "2020-08-04T22:08:00Z")
	firstTransaction := card.Transaction{
		Id:     1,
		Amount: 1203_91,
		Date:   firstDate,
		MCC:    "5812",
		Status: "Pending",
		Type:   card.WITHDRAW,
	}
	secondTransaction := card.Transaction{
		Id:     2,
		Amount: 2000_00,
		Date:   secondDate,
		MCC:    "",
		Status: "Done",
		Type:   card.DEPOSIT,
	}
	thirdTransaction := card.Transaction{
		Id:     3,
		Amount: 735_55,
		Date:   thirdDate,
		MCC:    "5411",
		Status: "Pending",
		Type:   card.WITHDRAW,
	}
	master := card.Card{
		Id:           1,
		Issuer:       "MasterCard",
		Balance:      100_000_00,
		Currency:     "RUB",
		Number:       "0001",
		Icon:         "https://...",
		Transactions: []card.Transaction{firstTransaction, secondTransaction, thirdTransaction},
	}
	fmt.Println(master)
	fmt.Println("Transactions len before add: ", len(master.Transactions))

	card.AddTransaction(&master, &card.Transaction{
		Id:     4,
		Amount: 321_00,
		Date:   time.Now(),
		MCC:    "5812",
		Status: "Done",
		Type:   card.WITHDRAW,
	})
	fmt.Println("Transactions len after added: ", len(master.Transactions))

	sumByMCC := card.SumByMCC(master.Transactions, []string{"5812"})

	fmt.Println("Sum by MCC: ", sumByMCC)

	fmt.Println("Exist code: ", card.TranslateMCC(master.Transactions[0].MCC))
	fmt.Println("Not exist code: ", card.TranslateMCC("2133"))

	fmt.Println("Last 2 transactions: ", card.LastNTransactions(&master, 2))
	fmt.Println("Last 2 transactions: ", len(card.LastNTransactions(&master, 2)))
	fmt.Println("Last 4 transactions because arg bigger then slice len: ", card.LastNTransactions(&master, 9))
	fmt.Println("Last 4 transactions because arg bigger then slice len: ", len(card.LastNTransactions(&master, 9)))
}
