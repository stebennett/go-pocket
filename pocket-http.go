package gopocket

import (
	"io/ioutil"
	"bytes"
	"encoding/json"
	"net/http"
)

func buildPocketRequest(pocketUrl string, body map[string]interface{}) (req* http.Request, err error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", pocketUrl, bytes.NewBuffer(bodyBytes))
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Accept", "application/json")
	return req, nil
}

func readPocketResponse(resp* http.Response, statusCheck int, value interface{}) (err error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, value)
	if err != nil {
		return err
	}

	return nil
}

