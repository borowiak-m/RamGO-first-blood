package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf
var typeOf = reflect.TypeOf

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) getBlanace() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not possible to withdraw")
	} else {
		pf("%d withdrawn Balance : %d\n", v, a.balance)
		a.balance -= v
	}
}

func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	var acct Account
	acct.balance = 100
	pl("Balance :", acct.getBlanace())
	for i := 0; i < 12; i++ {
		go acct.withdraw(10)
	}
	time.Sleep(2 * time.Second)

	useFunc(sumVals, 5, 8)
}
