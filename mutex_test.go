package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Counter : ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	balance float64
}

func (account *BankAccount) AddBalance(amount float64) {
	account.RWMutex.Lock()
	account.balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() float64 {
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance = u.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Azie",
		Balance: 10000,
	}
	user2 := UserBalance{
		Name:    "Melza",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1500)

	time.Sleep(1 * time.Second)

	fmt.Println("User: ", user1.Name, ", Balance: ", user1.Balance)
	fmt.Println("User: ", user2.Name, ", Balance: ", user2.Balance)
}
