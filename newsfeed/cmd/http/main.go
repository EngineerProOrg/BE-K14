package main

import (
	"fmt"

	"ep.k14/newsfeed/config"
	"ep.k14/newsfeed/internal/handler/http"
)

func main() {
	config, err := config.LoadHttpConfig()
	if err != nil {
		fmt.Println("err init config", err)
		return
	}
	fmt.Println("load config successfully", config)

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
