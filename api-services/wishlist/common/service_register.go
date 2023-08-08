package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func RegisDiscovery(name, port string) {

	url := os.Getenv("SERVER_IP") + "/register"
	type Payload struct {
		ServiceName string `json:"serviceName"`
		ServiceUrl  string `json:"serviceUrl"`
	}

	payload := Payload{
		ServiceName: name,
		ServiceUrl:  GetPublicIP() + ":" + port,
	}

	// Convert struct to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers if needed
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	// Read and display the response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println("Response Body:", buf.String())
}
