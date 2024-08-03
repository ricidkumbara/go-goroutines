package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestMutex(t *testing.T) {
	var x int = 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println(x)
}

func TestRWMutex(t *testing.T) {
	accont := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for J := 0; J < 100; J++ {
				accont.AddBalance(1)
				fmt.Println(accont.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(accont.GetBalance())
}
