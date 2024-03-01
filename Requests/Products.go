package requests

import (
	utility "CallPosAPI/Utility"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/go-resty/resty/v2"
)

func GetAllProducts(baseUrl string, accessToken string, count int) error {
	client := resty.New()
	var response []ProductResponse

	perPage := 100
	pages := int(math.Ceil(float64(count) / float64(perPage)))
	bar := utility.StartProgressBar(count)

	for i := 1; i <= pages; i++ {
		resp, err := client.R().
			SetAuthToken(accessToken).
			Get(fmt.Sprintf("%sproducts/%d/%d", baseUrl, i, perPage))
		if err != nil {
			return err
		}

		var jsonResponse []ProductResponse
		err = json.Unmarshal(resp.Body(), &jsonResponse)
		if err != nil {
			fmt.Println(err)
			return err
		}

		bar.UpdateStatus(len(jsonResponse))

		if len(jsonResponse) == 0 {
			break
		}

		response = append(response, jsonResponse...)
	}
	bar.Finish()

	println("\nProduct totall count " + strconv.Itoa(len(response)))

	duplicates := findDuplicates(response)
	if len(duplicates) > 0 {
		err := fmt.Errorf("Found duplicates: %v", duplicates)
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No duplicates found.")
	}

	return nil
}

func findDuplicates(nums []ProductResponse) []int {
	seen := make(map[int]bool)
	duplicates := make([]int, 0)

	for _, num := range nums {
		if seen[num.ID] {
			duplicates = append(duplicates, num.ID)
		} else {
			seen[num.ID] = true
		}
	}

	return duplicates
}

type ProductResponse struct {
	ID int `json:"id"`
}
