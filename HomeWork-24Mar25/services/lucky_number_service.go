package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateLuckyNumber() {
	// create a timer
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		//Ensure only 1 goroutine can read/write global variable LuckNumber at one time.
		Mu.Lock()
		LuckyNumber = rand.Intn(10) // ramdom number [0..9]
		fmt.Printf("\n [System start] generate lucky number: %d\n", LuckyNumber)
		Mu.Unlock()
		<-ticker.C // wait for 10 seconds then repeat the loop
	}
}
