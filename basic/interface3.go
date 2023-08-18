package basic

import "fmt"

type IConfig interface {
	Url() string
}

type Config struct {
	Host string
	Port int
}

func (c Config) Url() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func newConfig() IConfig {
	return Config{
		Host: "localhost",
		Port: 8080,
	}
}

func interface3() {
	config := newConfig()

	fmt.Printf("%#v\n", config.Url())
}
