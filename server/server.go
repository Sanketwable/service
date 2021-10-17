package server

import (
	"fmt"
	"service/config"
	"service/router"
)

func Run() {
	config.Load()
	fmt.Println("config file loaded")
	
	fmt.Println("DB loaded")
	fmt.Printf("\n\tListening.......[::]:%d \n", config.PORT)
	Listen(config.PORT)
}

func Listen(port int) {
	e := router.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
