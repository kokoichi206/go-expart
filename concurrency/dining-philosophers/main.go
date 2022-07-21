package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Each philosopher should eat at least 3 times.
const hunger = 3

var philosophers = []string{"John", "Plato", "Pascal", "Locke", "Socrates"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderFinished []string
// Fix race condition
var orderMutex sync.Mutex

func diningProblem(philosopher string, leftForl, rightFork *sync.Mutex) {
	defer wg.Done()

	fmt.Printf("%s is seated\n", philosopher)
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Printf("%s is hungry.\n", philosopher)
		time.Sleep(sleepTime)

		leftForl.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philosopher)

		fmt.Printf("%s has both forks, and is eating.\n", philosopher)
		time.Sleep(eatTime)

		// Give the philosopher some time ti think
		fmt.Printf("%s is thinking.\n", philosopher)
		time.Sleep(thinkTime)

		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork to his right.\n", philosopher)
		leftForl.Unlock()
		fmt.Printf("\t%s put down the fork to his left.\n", philosopher)
		time.Sleep(sleepTime)
	}

	fmt.Printf("%s is satisfied.\n", philosopher)
	time.Sleep(sleepTime)

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}

func main() {
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("--------------------------------")
	start := time.Now()

	wg.Add(len(philosophers))

	leftFork := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		rightFork := &sync.Mutex{}
		go diningProblem(philosophers[i], leftFork, rightFork)

		leftFork = rightFork
	}

	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("--------------------------------")
	fmt.Printf("Order finished: %s\n", strings.Join(orderFinished, ", "))
	fmt.Printf("It took %vms.\n", time.Since(start).Seconds())
}
