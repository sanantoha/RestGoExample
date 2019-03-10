package main

import (
	"time"
	"fmt"
)

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host string
	Port int
	Username string
	Password string
	Dbname string
}

type Config struct {
	DatabaseConfig DatabaseConfig `toml:"database"`
	ServerConfig ServerConfig `toml:"server"`
}

type User struct {
	Name string `json:"name,omitempty"`
	Age int 	`json:"age,omitempty"`
	LastUpdatetime time.Time `json:"last_updatetime,omitempty"`
}

func (user User) String() string {
	return fmt.Sprintf("User{name: %s, age: %d, last_updatetime: %s}", user.Name, user.Age, user.LastUpdatetime)
}

func (config Config) String() string {
	return fmt.Sprintf("Config{%s, %s}", config.ServerConfig, config.DatabaseConfig)
}

func (config ServerConfig) String() string {
	return fmt.Sprintf("ServerConfig{port: %d}", config.Port)
}

func (config DatabaseConfig) String() string {
	return fmt.Sprintf("DatabaseConfig{host: %s, port: %d, username: %s, password: ***, dbname: %s", config.Host, config.Port, config.Username, config.Dbname)
}