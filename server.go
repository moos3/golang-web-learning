package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/moos3/golang-web-learning/api"
	"github.com/moos3/golang-web-learning/models"
	"github.com/moos3/golang-web-learning/routes"
	"github.com/urfave/negroni"
	log "github.com/sirupsen/logrus"
)

// Config -
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Port string `json:"port"`
		User string `josn:"user"`

	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// LoadConfiguration -
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {
	if !os.Getenv("CONFIG_FILE") {
		ConfigFile = "config/config.development.json"
	} else {
		ConfigFile = os.Getenv("CONFIG_FILE")
	}
	config := LoadConfiguration(ConfigFile)
	fmt.Println(config.Database.Type)
	if config.Database.Type == 'sqlite' {
		db := models.NewSqliteDB(config.Database.Name)
	} else {
		db := models.NewDatabse(config.Database.Type, config.Database.Host, config.Database.Password, config.Database.Port, config.Database.Name, config.Database.User)
	}
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	if !Config.Port {
		Config.Port = "3000"
	}
	if !Config.Host {
		Config.Host = "0.0.0.0"
	}
	fmt.Println(Config.Host)
	n.Run(":3000")
}
