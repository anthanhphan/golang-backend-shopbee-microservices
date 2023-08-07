package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shopbee/common"
)

type userData struct {
	Data struct {
		Id     string `json:"id"`
		Status int    `json:"status"`
		Email  string `json:"email"`
		Role   string `json:"role"`
	} `json:"data"`
}

func getProfile(jwt string) (*common.SimpleUser, error) {
	var simpleUser common.SimpleUser

	endpoint := "http://13.54.238.78/api/v1/user/profile"
	jwtToken := jwt

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Decode and handle JSON response
	var responseData userData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	simpleUser.Status = responseData.Data.Status
	simpleUser.Email = responseData.Data.Email
	simpleUser.Role = responseData.Data.Role

	fmt.Print(simpleUser)
	return &simpleUser, nil
}
