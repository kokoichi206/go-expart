package main

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

func (c *Customer) UpdateAge(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		// deadlock が発生し panic する！！！
		return fmt.Errorf("age must be positive for customer %v", c)
	}

	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return fmt.Sprintf("Customer(id=%v, age=%v)", c.id, c.age)
}

func condSample() {
	type Donation struct {
		cond *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	f := func(goal int){
		donation.cond.L.Lock()
		for donation.balance < goal {
			// Wait の呼び出しはクリティカルセクション内で行う必要がある！
			donation.cond.Wait()
		}

		fmt.Printf("balance: %v\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go f(10)
	go f(20)

	for {
		time.Sleep(time.Second)

		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()

		// 条件が成立した（balance の更新）を通知する。
		donation.cond.Broadcast()
	}
}
