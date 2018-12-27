package etc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type (
	TiDB struct {
		Host 		string
		Port		int64
		Database	string
		User 		string
		Password 	string
	}

	Kafka struct {
		Hosts 	[]string
		Topic 	[]string
	}
	Redis struct {
		Sentinel 	[]string
		MasterName 	string
		PoolSize	int

		Host 		string
		Password 	string
		DB 			int
	}

	Config struct {
		Listen 	string
		Logger 	string
		KeyUrl 	string	`json:"KEY_API"`

		// TiDB 	*TiDB
		Kafka 	*Kafka
		Redis 	*Redis
	}
)

func (c *Config)LoadConfig(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New(fmt.Sprintf("load config %s failed. %s \n", path, err.Error()))
	}

	if len(content) == 0 {
		return errors.New("Config file is empty.")
	}

	if err := json.Unmarshal(content, c); err != nil {
		return errors.New(fmt.Sprintf("parse %s file error. ", path))
	}

	return nil
}