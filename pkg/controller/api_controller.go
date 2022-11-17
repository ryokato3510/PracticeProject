package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"PracticeProject/pkg/dto"

	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-password/password"
)

func Convert(i int) {
	var s string = strconv.Itoa(i)
	log.Info().Msg(s)
}

type ApiController interface {
	GetDoing(w http.ResponseWriter, r *http.Request)
	PostDoing(w http.ResponseWriter, r *http.Request)
}

type apiController struct {
}

func NewApiController() ApiController {
	return &apiController{}
}

func (ac *apiController) GetDoing(w http.ResponseWriter, r *http.Request) {

	var apiResponse dto.ApiResponse
	var message string = "Get Request reseived promptly."
	res, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		fmt.Printf("Error Can not generate password: %v", err)
	}
	fmt.Println(res)

	apiResponse.Content = message
	output, _ := json.MarshalIndent(apiResponse, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	// w.Write("Get Request reseived promptoly.")
}

func (ac *apiController) PostDoing(w http.ResponseWriter, r *http.Request) {

	var apiResponse dto.ApiResponse
	var message string = "Post Request reseived promptoly."

	apiResponse.Content = message
	output, _ := json.MarshalIndent(apiResponse, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	// w.Write("Get Request reseived promptoly.")
}
