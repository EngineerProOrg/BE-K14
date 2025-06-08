package main

import (
	"fmt"

	"ep.k14/newsfeed/internal/handler/http"
)

func main() {
	// create http server
	httpServer, err := http.New()
	if err != nil {
		fmt.Println("err init http server", err)
		return
	}

	err = httpServer.Start()
	if err != nil {
		fmt.Println("err start http server", err)
		return
	}

	// block here
}
