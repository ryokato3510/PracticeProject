package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello!")
	log.Info().Msg("hello world")
}
