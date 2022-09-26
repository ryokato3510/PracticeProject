package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"PracticeProject/pkg/controller"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var ac = controller.NewApiController()
var ro = controller.NewRouter(ac)

func main() {
	log.Info().Msg("hello world")
	loadEnv()

	config, err := LoadConfig("../../config/config.yml")
	if err != nil {
		fmt.Printf("Error Can not read config.yml: %v", err)
	}

	fmt.Printf("config: %v\n", config)

	controller.Convert(123)
	s := &http.Server{
		Addr:           ":8081",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/hello", ro.HandleTodosRequest)
	fmt.Println("Server Start")
	s.ListenAndServe()
	// log.Fatal()
}

// Read end file.
func loadEnv() {
	err := godotenv.Load("../../.env")

	if err != nil {
		fmt.Printf("Error Can not read .env file: %v", err)
	}

	message := os.Getenv("SAMPLE_MESSAGE")

	fmt.Println(message)
}

// Read config.yml
func LoadConfig(filePath string) (config *Config, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	result := &Config{}
	if err := yaml.Unmarshal(content, result); err != nil {
		return nil, err
	}

	return result, nil
}

type Config struct {
	Environment string `yaml:"environment"`

	SqlDriver string `yaml:"sqldriver"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Address   string `yaml:"address"`
	DataBase  string `yaml:"database"`
}
