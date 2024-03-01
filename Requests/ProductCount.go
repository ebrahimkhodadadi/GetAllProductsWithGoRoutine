package requests

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type RequestCount struct {
	Token string `json:"Token"`
}

type ResponseCount struct {
	IsValid      bool   `json:"isValid"`
	NewToken     string `json:"newToken"`
	ProductCount int    `json:"productCount"`
}

func GetProductCount(baseUrl string, accessToken string) (newAccessToken string, productCount int, err error) {
	fmt.Println("trying to get product count")
	client := resty.New()
	var response ResponseCount

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(RequestCount{Token: accessToken}).
		SetResult(&response).
		Post(baseUrl + "Users/refreshtoken")

	if err != nil || !response.IsValid {
		fmt.Println("GetProductCount status: ", resp.StatusCode())
		return "", 0, err
	}
	if resp.IsError() {
		return "", 0, fmt.Errorf("Error: %s", resp.Status())
	}

	fmt.Printf("product count is %v \n", response.ProductCount)

	if response.NewToken == "" {
		return accessToken, response.ProductCount, nil
	}

	return response.NewToken, response.ProductCount, nil
}
