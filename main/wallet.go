package main

import "fmt"

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}
