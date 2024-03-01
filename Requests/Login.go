package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type RequestLogin struct {
	UserName   string
	Password   string
	Url        string
	ShopId     int
	IsSubAdmin bool
}

type ResponseLogin struct {
	AccessToken string `json:"token"`
}

func Login(baseUrl string) (string, error) {
	fmt.Println("Try to login...")
	request, err := json.Marshal(RequestLogin{
		UserName: "username@gmail.com",
		Password: "123456789",
	})
	if err != nil {
		return "", err
	}

	response, err := http.Post(baseUrl+"Users/Authenticate", "application/json", bytes.NewBuffer(request))
	if err != nil {
		fmt.Printf("Status: %s\n", strconv.Itoa(response.StatusCode))
		return "", err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		responseText := string(responseData)
		fmt.Println(responseText)
		return "", err
	}
	var responseObject ResponseLogin
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("Login finished successfuly")
	//fmt.Println(responseObject.AccessToken)
	return responseObject.AccessToken, nil
}
