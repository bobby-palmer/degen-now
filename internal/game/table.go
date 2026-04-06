package game

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

type Table struct {
	*Config
}

func NewTable() *Table {
	return &Table{
		Config: NewConfig(),
	}
}
