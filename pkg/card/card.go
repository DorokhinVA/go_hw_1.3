package card

func Withdraw(card *Card, amount int64) {
	card.Balance -= amount
}

func Deposit(card *Card, amount int64) {
	card.Balance += amount
}

func Sum(cards []Card) int64 {
	total := int64(0)
	for _, card := range cards {
		total += card.Balance
	}
	return total
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func SumByMCC(transactions []Transaction, mcc []string) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		if contains(mcc, transaction.MCC) {
			result += transaction.Amount
		}
	}
	return result
}
func LastNTransactions(card *Card, n int) []Transaction {
	var resultLength int
	if len(card.Transactions) < n {
		resultLength = len(card.Transactions)
	} else {
		resultLength = n
	}

	tmp := make([]Transaction, resultLength)
	for index := 0; index < resultLength; index++ {
		tmp[index] = card.Transactions[len(card.Transactions)-index-1]
	}

	return tmp
}

func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
