package arrayswithgenerics

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from Account, to Account, amount float64) Transaction {
	return Transaction{from.Name, to.Name, amount}
}

func NewBalanceFor(acc Account, transactions []Transaction) Account {
	return Reduce(transactions, adjustBalance, acc)
}

func adjustBalance(acc Account, transaction Transaction) Account {
	if transaction.From == acc.Name {
		acc.Balance -= transaction.Sum
	}

	if transaction.To == acc.Name {
		acc.Balance += transaction.Sum
	}
	return acc
}
