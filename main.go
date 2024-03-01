package main

import (
	requests "CallPosAPI/Requests"
	"fmt"
	"time"
)

var baseUrl = "https://localhost:5001/api/v1.0/"

func main() {
	token, err := requests.Login(baseUrl)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for {
		newAccessToken, productCount, err := requests.GetProductCount(baseUrl, token)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		//err = requests.GetAllProducts(baseUrl, newAccessToken, productCount)
		err = requests.GetAllProductsWithGoRoutine(baseUrl, newAccessToken, productCount)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Stopping the program.")
			return
		}

		fmt.Println("Finish.")
		nextRunTime := time.Now().Add(time.Second * 60)
		fmt.Printf("Next run at: %s\n", nextRunTime)
		fmt.Println("-------------------------------------------------------------------")

		waitTime := time.Second * 60 // Adjust the waiting period as needed
		time.Sleep(waitTime)
	}
}
