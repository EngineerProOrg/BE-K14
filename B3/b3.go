package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bet struct {
	userID int
	amount int
}

type MoneyPool struct {
	bets        []Bet //store the bet history
	luckyNumber int
	totalMoney  int
}

var (
	lock sync.Mutex
	wg   sync.WaitGroup
)

func main() {
	moneyPool := &MoneyPool{}
	bet := make(chan Bet)
	luckyNumber := make(chan int)
	done := make(chan bool)

	go func() {
		//Generate lucky number goroutine
		for {
			select {
			case <-done:
				break
			}
			newNum := rand.Intn(10)
			fmt.Printf("System: %s: Lucky number changed to %d\n", time.Now().Format("2006-01-02 15:04:05"), newNum)
			luckyNumber <- newNum
			time.Sleep(10 * time.Second)
		}
	}()
	moneyPool.luckyNumber = <-luckyNumber
	fmt.Printf("(Start hệ thống generate lucky = %d)\n", moneyPool.luckyNumber)

	go func() {
		//Bet goroutine
		userID := 1
		for {
			var input string
			fmt.Printf(">User %d bet: \n", userID)
			fmt.Scanln(&input)

			if input == "END" {
				fmt.Printf("System: %s: Stopped input.\n", getTimestamp())
				done <- true
				break
			}
			var amount int
			_, err := fmt.Sscanf(input, "%d", &amount)
			if err != nil || amount < 1 || amount > 100 {
				fmt.Println("Invalid bet amount. Please enter a number between 1 and 100.")
				continue
			}
			bet <- Bet{userID: userID, amount: amount}
			userID++
		}
	}()

	exit := false
	for !exit {
		select {
		case num := <-luckyNumber:
			lock.Lock()
			moneyPool.luckyNumber = num
			lock.Unlock()
		case b := <-bet:
			lock.Lock()
			moneyPool.bets = append(moneyPool.bets, b)
			moneyPool.totalMoney += b.amount
			go findWinner(moneyPool)
			wg.Add(1)
			lock.Unlock()
		case <-done:
			exit = true
			fmt.Println("System: Exited.")
		}
	}
	close(bet)
	close(done)
	wg.Wait()
}

func getTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func findWinner(moneyPool *MoneyPool) {
	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()
	bet := moneyPool.bets[len(moneyPool.bets)-1]
	fmt.Printf("System: %s: User %d bet $%d, the current pool is $%d, waiting for %ds to receive result\n",
		time.Now().Format("2006-01-02 15:04:05"), bet.userID, bet.amount, moneyPool.totalMoney, bet.amount)
	time.Sleep(time.Duration(bet.amount) * time.Second)
	if bet.amount%10 == moneyPool.luckyNumber {
		fmt.Printf("System: %s: User %d hit the lucky number, get $%d\n", getTimestamp(), bet.userID, moneyPool.totalMoney)
		moneyPool.totalMoney = 0
		moneyPool.bets = make([]Bet, 0)
	} else {
		fmt.Printf("System: %s: Wish user %d lucky next time\n", getTimestamp(), bet.userID)
	}
}
