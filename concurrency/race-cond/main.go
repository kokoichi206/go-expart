package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Initial balance: %d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 500,
		},
		{
			Source: "Gifts",
			Amount: 10,
		},
		{
			Source: "Part time job",
			Amount: 50,
		},
		{
			Source: "Invenstments",
			Amount: 100,
		},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {

			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned %d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()
	fmt.Printf("Final bank balance: %d.00\n", bankBalance)
}

// var msg string
// var wg sync.WaitGroup

// pointer
// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()

// 	// Thread safe access to the data
// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "hi"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("hello", &mutex)
// 	go updateMessage("hello, guys", &mutex)
// 	wg.Wait()

// 	fmt.Println(msg)
// }
