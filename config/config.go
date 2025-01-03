package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Server struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	GrpcPort int    `yaml:"grpc_port"`
}

type Database struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

type Config struct {
	Server         Server   `yaml:"server"`
	MasterDb       Database `yaml:"master_db"`
	SlaveDb        Database `yaml:"slave_db"`
	SlaveDbEnable  bool     `yaml:"slave_db_enable"`
	EnableApiCache bool     `yaml:"enable_api_cache"`
	RPCs           []*RPC   `yaml:"rpcs"`
}

type RPC struct {
	RpcUrl           string   `yaml:"rpc_url"`
	ChainId          uint64   `yaml:"chain_id"`
	HeaderBufferSize uint64   `yaml:"header_buffer_size"`
	StartBlock       uint64   `yaml:"start_block"`
	EventStartBlock  uint64   `yaml:"event_start_block"`
	Contracts        []string `yaml:"contracts"`
}

func New(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
