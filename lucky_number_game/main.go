package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	ch    = make(chan bool)
	count int
	mu    sync.Mutex
)

func handleResult(numberUser int, duration int, bet int, luckyNumber int, count *int, pool *int, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(time.Duration(duration) * time.Second)
	defer timer.Stop()
	<-timer.C
	if bet%10 != luckyNumber {
		fmt.Printf("Wish user %d lucky next time\n", numberUser)
	} else {
		mu.Lock()
		fmt.Printf("User %d hit the lucky number, get %d$\n", numberUser, pool)
		*pool = 0
		mu.Unlock()
	}
	mu.Lock()
	*count--
	fmt.Printf("User remain: %d\n", *count)
	mu.Unlock()
	return
}

func inputData(numberUser *int, bet *string, pool *int, luckyNumber int, wg *sync.WaitGroup) {
	for {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("User %d bet: ", *numberUser)
		_, err := fmt.Scanf("%s", bet)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			return
		}
		//todo: out program
		if *bet == "END" || *bet == "END." {
			fmt.Printf("System: %s: Stopped input.\n", timeFormat)
			if count == 0 {
				ch <- true
				return
			}
		}
		betInt, err := strconv.Atoi(*bet)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		wg.Add(1)
		count++
		go handleResult(*numberUser, betInt, betInt, luckyNumber, &count, pool, wg)
		*pool += betInt
		fmt.Printf("System: %s: User %d bet $%s, the current pool is %d$, waiting for %ss to receive result\n", timeFormat, *numberUser, *bet, *pool, *bet)
		*numberUser++
	}
}

func endGame(ch chan bool) {
	select {
	case <-ch:
		fmt.Printf("System: Exited.")
		os.Exit(0)
	}
}

func updateValue(luckyNumber *int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().UnixNano())
			mu.Lock()
			*luckyNumber = rand.Intn(10)
			fmt.Println("\nUpdate lucky number: ", *luckyNumber)
			mu.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	luckyNumber := rand.Intn(10)
	numberUser := 1
	var bet string
	var pool int
	fmt.Println("Start game with lucky number: ", luckyNumber)
	go updateValue(&luckyNumber, &wg)
	go endGame(ch)
	inputData(&numberUser, &bet, &pool, luckyNumber, &wg)
	wg.Wait()
}
