package commands

type state struct (
	cfg *config.Config{}
)

type command struct (
	name string
	arguments []string
)

func handlerLogin(s *state, cmd command) error {
	
}