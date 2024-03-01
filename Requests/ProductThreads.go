package requests

import (
	"CallPosAPI/Utility"
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/go-resty/resty/v2"
)

var perPage = 100

func GetAllProductsWithGoRoutine(baseUrl string, accessToken string, count int) error {
	client := resty.New()
	var response []ProductResponse

	pages := int(math.Ceil(float64(count) / float64(perPage)))
	bar := Utility.StartProgressBar(count)

	var wg sync.WaitGroup
	maxWorkers := 10 // Adjust the number of workers as needed

	pageCh := make(chan int, pages)
	resultCh := make(chan []ProductResponse, pages)

	for i := 1; i <= pages; i++ {
		pageCh <- i
	}
	close(pageCh)

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(client, baseUrl, accessToken, pageCh, resultCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for jsonResponse := range resultCh {
		bar.UpdateStatus(len(jsonResponse))

		if len(jsonResponse) == 0 {
			break
		}

		response = append(response, jsonResponse...)
	}
	bar.Finish()

	fmt.Printf("\nProduct total count %d\n", len(response))

	duplicates := findDuplicates(response)
	if len(duplicates) > 0 {
		err := fmt.Errorf("Found duplicates: %v", duplicates)
		return err
	} else {
		fmt.Println("No duplicates found.")
	}

	return nil
}

func worker(client *resty.Client, baseUrl string, accessToken string, pageCh <-chan int, resultCh chan<- []ProductResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	for page := range pageCh {
		resp, err := client.R().
			SetAuthToken(accessToken).
			Get(fmt.Sprintf("%sproducts/%d/%d", baseUrl, page, perPage))
		if err != nil {
			fmt.Println(err)
			return
		}

		var jsonResponse []ProductResponse
		err = json.Unmarshal(resp.Body(), &jsonResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		resultCh <- jsonResponse
	}
}
