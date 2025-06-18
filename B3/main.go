package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type GameState struct {
	mu          *sync.Mutex
	pool        int
	luckyNumber int
	wg          *sync.WaitGroup
	isUserEnd   bool
}

func NewGameState() *GameState {
	return &GameState{
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}
}

func (g *GameState) updateLuckyNumber() {
	g.mu.Lock()
	g.luckyNumber = rand.Intn(10)
	g.mu.Unlock()
}

func (g *GameState) processUserBet(userId int, betAmount int) {
	g.mu.Lock()
	g.pool += betAmount
	g.mu.Unlock()

	timeStr := time.Now().Format("02-Jan-2006 15:04:05")
	fmt.Printf("System: %s User %d bet %d$, the current pool is %d$, waiting for %ds to receive result\n",
		timeStr, userId, betAmount, g.pool, betAmount)
}

func (g *GameState) checkWinner(userId, betAmount int) {
	timeStr := time.Now().Format("02-Jan-2006 15:04:05")
	g.mu.Lock()
	defer g.mu.Unlock()

	if betAmount%10 == g.luckyNumber {
		fmt.Printf("\nSystem: %s User %d hit the lucky number, get %d\n", timeStr, userId, g.pool)
		g.pool = 0
	} else {
		fmt.Printf("\nSystem: %s Wish User %d lucky next time\n", timeStr, userId)
	}
}

func startLuckyNumberGenerator(game *GameState) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		game.updateLuckyNumber()
	}
}

func handleUserInput(game *GameState, userId int) (int, bool) {
	fmt.Printf("User %d bet: ", userId)
	var inputStr string
	_, err := fmt.Scan(&inputStr)
	if err != nil {
		fmt.Printf("Invalid input\n")
		return 0, false
	}

	timeStr := time.Now().Format("02-Jan-2006 15:04:05")
	if inputStr == "END" {
		game.mu.Lock()
		game.isUserEnd = true
		game.mu.Unlock()
		fmt.Printf("System: %s Stopped input\n", timeStr)
		return 0, true
	}

	betAmount, err := strconv.Atoi(inputStr)
	if err != nil || betAmount < 0 || betAmount > 9 {
		fmt.Printf("Invalid input. Please enter a number in range [0, 9] or 'END'\n")
		return 0, false
	}
	return betAmount, false
}

func processGameRound(game *GameState, userId, betAmount int) {
	game.wg.Add(1) 
	game.processUserBet(userId, betAmount)

	time.AfterFunc(time.Duration(betAmount)*time.Second, func() {
		game.checkWinner(userId, betAmount)
		game.wg.Done() 
	})
}

func runGameLoop(game *GameState, exitChannel chan struct{}) {
	userId := 1
	go func() {
		for {
			betAmount, shouldExit := handleUserInput(game, userId)
			if shouldExit {
				go func() {
					game.wg.Wait()
					close(exitChannel)
				}()
				return
			}

			if betAmount > 0 {
				processGameRound(game, userId, betAmount)
				userId++
			}
		}
	}()
}

func main() {
	game := NewGameState()
	exitChannel := make(chan struct{})

	// Start lucky number generator
	go startLuckyNumberGenerator(game)

	// Start game loop
	runGameLoop(game, exitChannel)

	// Wait for game to end
	<-exitChannel
}
