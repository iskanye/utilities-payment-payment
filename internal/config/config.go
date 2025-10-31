package config

type Config struct {
	Host    string   `yaml:"host" env-default:"localhost"`
	Port    int      `yaml:"port" env-required:"true"`
	Billing HostPort `yaml:"billing"`
}

type HostPort struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-required:"true"`
}
