package main

import (
	"./config"
	"fmt"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	fmt.Println(config.App.DbHost)
	fmt.Println(config.App.DbUser)
	fmt.Println(config.App.DbPassword)
}
