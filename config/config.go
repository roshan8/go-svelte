package config

// import (
// 	"gopkg.in/yaml.v3"
// )

// var Env config

// type config struct {
// 	Db dbConfig `yaml:"db"`
// }

// type dbConfig struct {
// 	Port     int    `yaml:"port"`
// 	Host     string `yaml:"host"`
// 	User     string `yaml:"user"`
// 	Password string `yaml:"password"`
// 	SslMode  string `yaml:"sslmode"`
// 	Dbname   string `yaml:"dbname"`
// }

// func Load() error {
// 	var config config

// 	err := yaml.Unmarshal(Config, &config)
// 	if err != nil {
// 		return err
// 	}

// 	Env = config

// 	return nil
// }
