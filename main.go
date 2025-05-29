package main

/**
B3: Mô phỏng 1 cái game lucky number đa luồng như sau.
Lucky number là 1 số ngẫu nhiên trong khoảng [0, 9], được random bởi hệ thống mỗi 10s.
Mọi người sẽ tạo 1 cái chương trình cho phép người dùng nhập số tiền bet từ màn hình (giới hạn 1-100).
Mỗi lần người dùng bet số tiền x:
Cộng số tiền đó vào pool tiền thưởng,
In ra màn hình: System: YYYY-MM-DD HH:mm:ss: User {{i}} bet ${{x}}, the current pool is {{pool+=x}}$, waiting for {{x}}s to receive result,
Quay xổ số: Người dùng sẽ phải chờ x giây để nhận kết quả. Sau x giây, so sánh nếu x % 10 = lucky number thì thông báo người dùng trúng thưởng, trả hết số tiền trong pool hiện tại, nếu ko trúng, thông báo ko trúng.,
In ra màn hình kết quả:
System: YYYY-MM-DD HH:mm:ss: User {{i}} hit the lucky number, get {{x}}$ (nếu trúng)
System: YYYY-MM-DD HH:mm:ss: Wish user {{i}} lucky next time (nếu fail)
Trong lúc người dùng thứ i chờ kết quả, vẫn tiếp tục cho phép người dùng tiếp theo nhận số tiền.,
Khi có 1 người nhập "END", dừng cho phép nhập input, chờ cho các lượt quay xổ số trước đó kết thúc và thông báo kết quả
Ví dụ:
```
(Start hệ thống generate lucky = 2)
>User 1 bet: 5
System: 2021-03-21 00:00:00: User 1 bet $5, the current pool is 5$, waiting for 5s to receive result
>User 2 bet: 3
System: 2021-03-21 00:00:01: User 2 bet 3$, the current pool is 8$, waiting for 3s to receive result
>User 3 bet: 1
System: 2021-03-21 00:00:02: User 3 bet: 1$, the current pool is 9$, waiting for 1s to receive result
>User 4 bet: END.
System: 2021-03-21 00:00:02: Stopped input.
System: 2021-03-21 00:00:03: Wish user 3 lucky next time
System: 2021-03-21 00:00:04: User 2 hit the lucky number, get 9$
System: 2021-03-21 00:00:05: Wish user 1 lucky next time
System: Exited.
```
*/

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Global game state with thread-safe access
type GameState struct {
	prizePool   int
	luckyNumber int
	userCounter int
	poolMutex   sync.Mutex
	luckyMutex  sync.Mutex
	userMutex   sync.Mutex
	shutdown    chan bool
	wg          sync.WaitGroup
}

func NewGameState() *GameState {
	return &GameState{
		prizePool:   0,
		luckyNumber: rand.Intn(10),
		userCounter: 0,
		shutdown:    make(chan bool),
	}
}

func getTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Thread-safe methods for GameState
func (gs *GameState) addToPrizePool(amount int) int {
	gs.poolMutex.Lock()
	defer gs.poolMutex.Unlock()
	gs.prizePool += amount
	return gs.prizePool
}

func (gs *GameState) takePrizePool() int {
	gs.poolMutex.Lock()
	defer gs.poolMutex.Unlock()
	pool := gs.prizePool
	gs.prizePool = 0
	return pool
}

func (gs *GameState) setLuckyNumber(num int) {
	gs.luckyMutex.Lock()
	defer gs.luckyMutex.Unlock()
	gs.luckyNumber = num
}

func (gs *GameState) getLuckyNumber() int {
	gs.luckyMutex.Lock()
	defer gs.luckyMutex.Unlock()
	return gs.luckyNumber
}

func (gs *GameState) getNextUserId() int {
	gs.userMutex.Lock()
	defer gs.userMutex.Unlock()
	gs.userCounter++
	return gs.userCounter
}

// Lucky number generator - runs every 10 seconds
func (gs *GameState) luckyNumberGenerator() {
	defer gs.wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-gs.shutdown:
			return
		case <-ticker.C:
			newLucky := rand.Intn(10)
			gs.setLuckyNumber(newLucky)
			fmt.Printf("System: %s: New lucky number generated: %d\n", getTimestamp(), newLucky)
		}
	}
}

// Process individual bet - runs as separate goroutine for each bet
func (gs *GameState) processBet(userId, betAmount int) {
	defer gs.wg.Done()

	// Add bet to prize pool
	currentPool := gs.addToPrizePool(betAmount)

	// Log bet placement
	fmt.Printf("System: %s: User %d bet $%d, the current pool is %d$, waiting for %ds to receive result\n",
		getTimestamp(), userId, betAmount, currentPool, betAmount)

	// Wait for bet amount seconds
	time.Sleep(time.Duration(betAmount) * time.Second)

	// Check win condition
	currentLucky := gs.getLuckyNumber()
	if betAmount%10 == currentLucky {
		// Winner - take entire pool
		wonAmount := gs.takePrizePool()
		fmt.Printf("System: %s: User %d hit the lucky number, get %d$\n",
			getTimestamp(), userId, wonAmount)
	} else {
		// Loser
		fmt.Printf("System: %s: Wish user %d lucky next time\n",
			getTimestamp(), userId)
	}
}

// Handle user input
func (gs *GameState) handleUserInput() {
	defer gs.wg.Done()

	for {
		userId := gs.getNextUserId()
		var input string
		fmt.Printf("> User %d bet: ", userId)
		_, err := fmt.Scanln(&input)
		if err != nil {
			continue
		}

		if input == "END" {
			fmt.Printf("System: %s: Stopped input.\n", getTimestamp())
			close(gs.shutdown)
			return
		}

		// Parse bet amount
		amount, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number or 'END'.")
			gs.userMutex.Lock()
			gs.userCounter-- // Revert user counter for invalid input
			gs.userMutex.Unlock()
			continue
		}

		// Validate bet range
		if amount < 1 || amount > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			gs.userMutex.Lock()
			gs.userCounter-- // Revert user counter for invalid input
			gs.userMutex.Unlock()
			continue
		}

		// Start bet processing in separate goroutine
		gs.wg.Add(1)
		go gs.processBet(userId, amount)
	}
}

func main() {
	fmt.Printf("System: %s: Started.\n", getTimestamp())

	// Create game state
	gs := NewGameState()

	// Display initial lucky number
	fmt.Printf("System: %s: Initial lucky number: %d\n", getTimestamp(), gs.getLuckyNumber())

	// Start lucky number generator
	gs.wg.Add(1)
	go gs.luckyNumberGenerator()

	// Start user input handler
	gs.wg.Add(1)
	go gs.handleUserInput()

	// Wait for all goroutines to complete
	gs.wg.Wait()

	fmt.Printf("System: %s: Exited.\n", getTimestamp())
}
