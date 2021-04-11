package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "./config.yml"
)

type serverConfig struct {
	Port string `yaml:"port"`
}

func (s *serverConfig) getConfig(f string) error {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		fmt.Println("No configuration file found. Using default configuration.")
		s.Port = ":80"
		return nil
	}

	fBytes, err := os.ReadFile("./config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fBytes, s)
	if err != nil {
		return err
	}
	return nil
} // end getConfig

func main() {

	config := serverConfig{}
	err := config.getConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./index.html")
	})
	fmt.Printf("Server listening on port %v\n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
