package main

import (
	"awesomeProject/internal/app"
	"flag"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(" ------------------------------------------ ")
	fmt.Println()
	app.Run(ParseFlags())
}

func ParseFlags() string {
	var configPath string

	flag.StringVar(&configPath, "config", "configs/appConfig.yml", "path to config file")

	flag.Parse()

	return configPath
}
