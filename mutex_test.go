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

/* DEADLOCK
Deadlock terjadi ketika dua goroutine yang sedang berjalan terhubung dan terdeteksi sebagai deadlock.
Berikut Simulasi Deadlock (1)
*/

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second) // Simulasi proses transfer

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Albany",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Asep",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance = ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance = ", user2.Balance)
}
