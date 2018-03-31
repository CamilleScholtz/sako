package main

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

var config Config

// Config is a stuct with all config values. See `runtime/config.toml` for more
// information about these values.
type Config struct {
	Daemon string
	RPC    string

	Host string

	Currency string
}

// parseConfig parses a toml config.
func parseConfig() error {
	//d, err := homedir.Dir()
	//if err != nil {
	//	return err
	//}

	d, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(path.Join(d, "runtime", "config.toml"),
		&config); err != nil {
		return err
	}

	return nil
}
