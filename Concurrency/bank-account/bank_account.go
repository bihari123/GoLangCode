package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	status  bool
	sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	var acc = Account{balance: amount, status: true}
	return &acc
	panic("Please implement the Open function")
}

func (a *Account) Balance() (int64, bool) {
	if a.status {
		return a.balance, true
	}
	return 0, false
	panic("Please implement the Balance function")
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.status {
		a.Lock()
		defer a.Unlock()
		if amount < 0 {
			if a.balance+amount < 0 {
				return 0, false
			}
		}
		a.balance += amount

	} else {
		return 0, false
	}

	return a.balance, true

	panic("Please implement the Deposit function")
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.status {
		a.status = false
		return a.balance, true
	}
	return 0, false

	panic("Please implement the Close function")
}
