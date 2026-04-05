package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RwMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RwMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {

	account.RwMutex.RLock()
	balance := account.Balance
	account.RwMutex.RUnlock()

	return balance
}

func TestRWMuutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {

				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

// DeadLock
type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() { user.Mutex.Lock() }

func (user *UserBalance) Unlock() { user.Mutex.Unlock() }

func (user *UserBalance) Change(amount int) { user.Balance = user.Balance + amount }

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {

	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()

}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bani",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Heri",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000) // Deadlock
	go Transfer(&user2, &user1, 200000) // Deadlock

	time.Sleep(10 * time.Second)

	fmt.Println("User:", user1.Name, "Balance:", user1.Balance)
	fmt.Println("User:", user2.Name, "Balance:", user2.Balance)
}
