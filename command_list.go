package main

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {

}

func (c *commands) run(s *state, cmd command) error {
	return nil
}
