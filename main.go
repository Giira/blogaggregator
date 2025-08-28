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
	s := state{}
	s.cfg = &cfg
}
