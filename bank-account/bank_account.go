package account

import "sync"

// Define the Account type here.
type Account struct {
	AccMux        sync.Mutex
	BalanceAmount int64
	IsClosed      bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{BalanceAmount: amount}
}

func (a *Account) Withdraw(amount int64) (int64, bool) {
	if a.IsClosed || amount > a.BalanceAmount {
		return a.BalanceAmount, false
	}
	a.AccMux.Lock()
	defer a.AccMux.Unlock()
	if a.IsClosed || amount > a.BalanceAmount {
		return a.BalanceAmount, false
	}
	a.BalanceAmount -= amount
	return a.BalanceAmount, true
}

func (a *Account) Balance() (int64, bool) {
	if a.IsClosed {
		return a.BalanceAmount, false
	}
	a.AccMux.Lock()
	defer a.AccMux.Unlock()
	if a.IsClosed {
		return a.BalanceAmount, false
	}
	return a.BalanceAmount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if amount < 0 {
		return a.Withdraw(amount * -1)
	}
	if a.IsClosed {
		return a.BalanceAmount, false
	}
	a.AccMux.Lock()
	defer a.AccMux.Unlock()
	if a.IsClosed {
		return a.BalanceAmount, false
	}
	a.BalanceAmount += amount
	return a.BalanceAmount, true
}

func (a *Account) Close() (int64, bool) {

	if a.IsClosed {
		return a.BalanceAmount, false
	}
	a.AccMux.Lock()
	defer a.AccMux.Unlock()
	if a.IsClosed {
		return a.BalanceAmount, false
	} else {
		payout := a.BalanceAmount
		a.IsClosed = true
		a.BalanceAmount = 0
		return payout, a.IsClosed
	}
}
