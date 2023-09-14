package config

import (
	"fingerGuessWeb3GameGoBackend/pkg/global"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("config")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Config{vp}, nil
}
func (c *Config) ReadConfig(key string, value interface{}) error {
	err := c.vp.UnmarshalKey(key, value)
	if err != nil {
		return err
	}
	return nil
}

func SetupConfig() {
	config, err := NewConfig()
	if err != nil {
		log.Panic("SetupConfig - NewConfig ：", err)
	}
	err = config.ReadConfig("DataBase", &global.DbConfig)
	if err != nil {
		log.Panic("SetupConfig - DataBase error ： ", err)
	}
	err = config.ReadConfig("BlockChain", &global.BlockConfig)
	if err != nil {
		log.Panic("SetupConfig - BlockChain error ： ", err)
	}
}
func SetupDbEngine() {
	var err error
	global.DbEngine, err = NewDbEngine(global.DbConfig)
	if err != nil {
		log.Panic("SetupConfig - DataBase error ： ", err)
	}
}
func SetupTable() {
	err := Migrate()
	if err != nil {
		log.Panic("SetupTable - Migrate error ： ", err)
	}
}
func SetupEthClient() {
	var err error
	global.RpcClient, err = NewEthClient()
	if err != nil {
		log.Panic("SetupEthClient - NewEthClient error ： ", err)
	}
}
