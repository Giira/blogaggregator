package main

import (
	"fmt"

	"github.com/Giira/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	cfg.SetUser("Euan")
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Current user: %s\nConfig database url: %s\n", cfg.Current_user_name, cfg.Db_url)
}
