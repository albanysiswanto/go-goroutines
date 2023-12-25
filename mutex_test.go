package belajar_go_goroutines

/*
SOLUSI RACE CONDITION!
Dengan menggunakan Library Mutex, kita bisa menghindari race condition.

KAPAN HARUS MENGGUNAKAN MUTEX?
Ketika kita memiliki variable sharing(yang di akses beberapa goroutine) di aplikasi kita, kita harus menggunakan mutex.
*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 100; j++ {
			mutex.Lock()
			x += 1
			mutex.Unlock()
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

// RWMutex (Read-Write Mutex) = Digunakan untuk menghindari race condition
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

// Write
func (account *BankAccount) addBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

// Read
func (account *BankAccount) getBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			account.addBalance(1)
			fmt.Println(account.getBalance())
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance = ", account.getBalance())
}
