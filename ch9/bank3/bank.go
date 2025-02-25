package bank

import "sync"


var (
	mu  sync.Mutex
	balance int
)


func Deposit(amount int){
	mu.Lock()
	balance += amount
	mu.Unlock()
}


func Balance() int{
	mu.Lock()
	t := balance
	mu.Unlock()
	return t
}