package bank

import "sync"

var (
	mu      sync.RWMutex
	balance int
)

// 要求调用deposit的函数已经获取了互斥锁
func deposit(amount int) {
	balance += amount
}

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func WithDraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)

	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
