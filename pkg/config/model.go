package config

type Config struct {
	Server *Server `yaml:"server"`
	Auth   *Auth   `yaml:"auth"`
	Gin    *Gin    `yaml:"gin"`
	Log    *Log    `yaml:"log"`
}

type Server struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
}

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Gin struct {
	Mode string `yaml:"mode"`
}

type Log struct {
	Level string `yaml:"level"`
}
