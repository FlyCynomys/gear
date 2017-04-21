package conf

import "encoding/json"
import "errors"
import "io/ioutil"

var cfg *Config

func Init(filapth string) {
	cfg = new(Config)
	loadConfig(filapth)
}

func GetCfg() *Config {
	if cfg == nil {
		return nil
	}
	return cfg
}

type Config struct {
	DbAddress    string `json:"dbaddress"`
	DbPassword   string `json:"dbpassword"`
	DbUser       string `json:"dbuser"`
	RedisUse     bool   `json:"redis"`
	RedisAddress string `json:"redisaddress"`
	Port         string `json:"port"`
}

func (c *Config) FormatFromJson(data []byte) (bool, error) {
	if data == nil {
		return false, errors.New("data is nil")
	}
	json.Unmarshal(data, c)
	return true, nil
}

func loadConfig(filepath string) (bool, error) {
	if filepath == "" {
		return false, errors.New("config file is empty")
	}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false, err
	}
	_, err = cfg.FormatFromJson(data)
	if err != nil {
		return false, err
	}
	return true, nil
}
