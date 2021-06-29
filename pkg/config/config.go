package config

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml"
)

//Config is the struct to save and load the config file
type Config struct {
	Telegram TelegramConf
	Database DatabaseConf
}

//TelegramConf holds the token for the bot and the timeout for the poller
type TelegramConf struct {
	Token         string
	PollerTimeout int
}

//DatabseConf holds the Driver and the Connection string
type DatabaseConf struct {
	Driver     string
	Connection string
}

//LoadConfig loads the config from the filepath and gives back a Config or an error
func LoadConfig(filepath string) (Config, error) {
	var res Config
	file, err := os.Open(filepath)
	if err != nil {
		return res, errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()
	decoder := toml.NewDecoder(file)
	err = decoder.Decode(&res)
	if err != nil {
		return res, errors.New("Error decoding file: " + err.Error())
	}
	return res, nil

}

//SaveConfig save a config file to the filepath
func (c *Config) SaveConfig(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return errors.New("Error creating file: " + err.Error())
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return errors.New("Error encoding file: " + err.Error())
	}
	return nil
}
