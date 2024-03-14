package database

type Config struct {
	Driver string `yaml:"driver"`

	IP   string `yaml:"ip"`
	Port string `yaml:"port"`

	Database string `yaml:"database"`
	ID       string `yaml:"id"`
	Password string `yaml:"password"`
}
