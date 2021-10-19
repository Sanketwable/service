package server

import (
	"fmt"
	"os"
	"service/config"
	"service/router"
)

func Run() {
	config.Load()
	fmt.Println("config file loaded")
	
	fmt.Println("DB loaded")
	fmt.Printf("\n\tListening.......[::]:%s \n", os.Getenv("PORT") )
	Listen(os.Getenv("PORT"))
}

func Listen(port string) {
	e := router.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
