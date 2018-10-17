package commands

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/cpuguy83/strongerrors"
	"github.com/pkg/errors"
)

// UserConfig represents the user configuration read from a config file
type UserConfig struct {
	Subscription string
	Location     string

	Profile struct {
		Kubernetes struct {
			Version       string
			NetworkPolicy string
			NetworkPlugin string
		}
		Leader struct {
			Linux struct {
				Distro string
				SKU    string
				Count  *int
			}
		}
		Agent struct {
			Linux struct {
				Distro              string
				SKU                 string
				Count               *int
				AvailabilityProfile string
			}
			Windows struct {
				SKU                 string
				Count               *int
				AvailabilityProfile string
			}
		}
		Auth struct {
			Linux struct {
				User          string
				PublicKeyFile string
			}
			Windows struct {
				User         string
				PasswordFile string
			}
		}
	}
}

// ReadUserConfig reads the config from the provided path
// If
func ReadUserConfig(configPath string) (UserConfig, error) {
	var cfg UserConfig

	f, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, strongerrors.NotFound(errors.Wrap(err, "user config file not found"))
		}
		return cfg, errors.Wrap(err, "could not open specified config file path")
	}

	if _, err := toml.DecodeReader(f, &cfg); err != nil {
		return cfg, errors.Wrap(err, "error decoding user config")
	}
	return cfg, nil
}
