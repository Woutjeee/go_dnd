package configuration

type CommandExecuter interface {
	Execute(cfg *Config, flags map[string]string)
}

type Config struct {
	LastCommand CommandExecuter
	LastFlags   map[string]string
}
