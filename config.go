package main

import (
	"os"
	"path"
	"regexp"

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
	l, err := locateConfig()
	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(l, &config); err != nil {
		return err
	}

	return nil
}

func locateConfig() (string, error) {
	//d, err := homedir.Dir()
	//if err != nil {
	//	return "", err
	//}

	d, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(d, "runtime", "config.toml"), nil
}

func modifyConfig(f []byte, k, v string) []byte {
	re := regexp.MustCompile("(?m)(^[[:space:]]*" + k +
		"[[:space:]]*=[[:space:]]*\")[[:upper:]]*(\".*)")
	f = re.ReplaceAll(f, []byte("${1}"+v+"${2}"))

	return f
}
