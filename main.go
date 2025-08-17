package blogaggregator

import (
	"fmt"

	"github.com/Giira/blogaggregator/internal/config"
)

func main() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}
	cfg.SetUser("Euan")
	config.Read()
	fmt.Printf("Current user: %s\nConfig database url: %s\n", cfg.Current_user_name, cfg.Db_url)
	return nil
}
