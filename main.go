package main

import (
	"flag"
	"os"
)

var (
	version    = "v0.1.0"
	configFile = flag.String("configFile", "config.json", "config file name")
)

type Config struct {
	Local  string `json:"local"`
	Route  string `json:"route"`
	Remote string `json:"remote"`
}

func loadConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var config Config

	f.Read([]byte("local"))

}

func main() {
	flag.Parse()

}
