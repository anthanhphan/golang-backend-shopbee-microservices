package common

import (
	"encoding/json"
	"net/http"
)

func GetPublicIP() string {
	type IPResponse struct {
		IP string `json:"ip"`
	}

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	var ipResponse IPResponse
	err = json.NewDecoder(resp.Body).Decode(&ipResponse)
	if err != nil {
		return ""
	}

	return "http://" + ipResponse.IP
}
