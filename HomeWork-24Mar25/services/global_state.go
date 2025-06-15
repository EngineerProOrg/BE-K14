package services

import "sync"

var (
	LuckyNumber int
	Pool        int
	Mu          sync.Mutex
	Wg          sync.WaitGroup
	DoneChan    = make(chan bool)
	TotalInputs int
)
