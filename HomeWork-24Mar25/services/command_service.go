package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func HandleUserInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command (e.g., User 1 bet 10 or END):")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if strings.ToUpper(input) == "END" {
			fmt.Println("Program terminated. Waiting for all users to finish...")
			Wg.Wait() // Wait until all users complete bet
			fmt.Println("All bets processed. Exiting.")
			break
		}

		// Split from input
		parts := strings.Split(input, " ")
		if len(parts) != 4 || strings.ToLower(parts[0]) != "user" || strings.ToLower(parts[2]) != "bet" {
			fmt.Println("❌ Invalid format. Please use: User {id} bet {amount}")
			continue
		}

		userID := parts[1]
		betStr := parts[3]

		// Convert bet to integer number
		bet, err := strconv.Atoi(betStr)
		if err != nil || bet < 1 || bet > 100 {
			fmt.Println("❌ Bet amount must be between 1 and 100.")
			continue
		}

		// ✅ Valid bet
		fmt.Printf("✅ Accepted bet of %d$ from User %s\n", bet, userID)

		// Spawn a goroutine to handle bet
		Wg.Add(1)
		go func() {
			defer Wg.Done()
			handleSingleBet(userID, bet)
		}()
	}
}

func handleSingleBet(userId string, bet int) {
	fmt.Printf("[System] User %s bet %d$, the current pool is %d$, waiting for %ds to receive result\n", userId, bet, Pool+bet, bet)

	// add bet money to pool
	Mu.Lock()
	Pool += bet
	Mu.Unlock()

	time.Sleep(time.Duration(bet) * time.Second)

	Mu.Lock()
	currentLucky := LuckyNumber
	currentPool := Pool
	Mu.Unlock()

	if bet%10 == currentLucky {
		// winer user
		Mu.Lock()
		Pool = 0 // reset pool
		Mu.Unlock()

		fmt.Printf("[System] User %s hit the lucky number, get %d$\n", userId, currentPool)
	} else {
		//
		fmt.Printf("[System] Wish user %s lucky next time\n", userId)
	}
}
