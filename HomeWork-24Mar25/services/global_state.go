package services

import "sync"

var (
	LuckyNumber int
	Pool        int
	Mu          sync.Mutex
	Wg          sync.WaitGroup
)
