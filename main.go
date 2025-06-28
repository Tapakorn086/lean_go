package main

import (
	"fmt"
	"learning/config"

	"github.com/roonglit/credentials/pkg/credentials"
)

func main() {
	reader := credentials.NewConfigReader()

	var config config.Config
	if err := reader.Read("debug", &config); err != nil {
		panic(err)
	}

	fmt.Println("App Name:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("DBURI:", config.DBURI)

	if err := reader.Read("release", &config); err != nil {
		panic(err)
	}

	fmt.Println("App Name:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("DBURI:", config.DBURI)
}
