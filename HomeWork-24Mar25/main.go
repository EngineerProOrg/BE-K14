package main

import "HomeWork-24Mar25/services"

/*
func main() {
	//go services.GenerateLuckyNumber()

	//services.HandleUserInput()

	// Neu khong dung go routine cac ham duoi day se chay tuan tu
	// services.HandleUserBet("user 1", 5)
	// services.HandleUserBet("user 2", 3)
	// services.HandleUserBet("user 3", 2)


	// Phai dung go routine cac ham duoi day xu ly doc lap ko can doi nhau
	go services.HandleUserBet("user 1", 5)
	go services.HandleUserBet("user 2", 3)
	go services.HandleUserBet("user 3", 12)

	time.Sleep(2 * time.Second)
}*/

func main() {
	go services.GenerateLuckyNumber()

	services.HandleUserInput()
}
