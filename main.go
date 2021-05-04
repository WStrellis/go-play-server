package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	DEFAULT_CONFIG_FILE = "./config.yml"
)

type serverConfig struct {
	Port string `yaml:"port"`
}

func (s *serverConfig) readConfigFile(f string) error {
	fBytes, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fBytes, s)
	if err != nil {
		return err
	}
	return nil

} // end readConfig

func (s *serverConfig) readEnv() {
	fmt.Println("Reading from environment variables")
	s.Port = os.Getenv("SERVER_PORT")
} // end readEnv

func (s *serverConfig) useDefaults() {
	fmt.Println("Using defaults")
	if s.Port == "" {
		s.Port = ":80"
	}
}

func (s *serverConfig) getConfig(f string) error {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		fmt.Printf("%v not found.\n", f)
		s.readEnv()
		s.useDefaults()
		return nil
	}

	err := s.readConfigFile(f)
	if err != nil {
		return err
	}
	return nil
} // end getConfig

func main() {

	config := serverConfig{}
	configFile := os.Getenv("SERVER_CONFIG_FILE")
	if configFile == "" {
		fmt.Printf("SERVER_CONFIG_FILE not set. Using default %v\n", DEFAULT_CONFIG_FILE)
		configFile = DEFAULT_CONFIG_FILE
	}
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
