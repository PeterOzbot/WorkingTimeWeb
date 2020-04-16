package config

import (
	"log"
	"workingtimeweb/server/core"

	"github.com/BurntSushi/toml"
	"github.com/matryer/resync"
)

var (
	configuration *core.Config
	once          resync.Once
)

// Configuration : Returns configuration and if it was not yet loaded it reads configuration from file.
func Configuration() *core.Config {
	once.Do(func() {

		// read config
		tomlData := readConfig()

		// decode it
		if _, err := toml.Decode(tomlData, &configuration); err != nil {
			log.Fatal("Failed decoding config file.")
		}
	})

	return configuration
}

func readConfig() string {

	byteData, err := iouUtil.ReadFile("../environment.toml")
	if err != nil {
		log.Fatal(err)
	}

	if byteData == nil || len(byteData) == 0 {
		log.Fatal("Failed loading config file: environment.toml")
	}

	return string(byteData)
}
