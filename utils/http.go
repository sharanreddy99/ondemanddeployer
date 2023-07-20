package utils

import (
	"encoding/json"
	"io"

	"net/http"
)

func ExecuteRequest(req *http.Request) interface{} {
	var respMap interface{}
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		Log("Error occured while executing external API: ", err.Error())
		return respMap
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Log("Error occured while reading body of external API: ", err.Error())
		return respMap
	}

	err = json.Unmarshal(body, &respMap)
	if err != nil {
		Log("Error occured while unmarshalling body of external API: ", err.Error())
		return respMap
	}

	return respMap
}
