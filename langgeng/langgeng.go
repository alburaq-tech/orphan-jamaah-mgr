package langgeng

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	apiHost = os.Getenv("LANGGENG_URL")
	apiKey  = os.Getenv("LANGGENG_APIKEY")
)

func CallLanggengSubscribeCampaign(campaign_code, name, phone, customData string) (map[string]interface{}, error) {
	hc := http.Client{}

	form := url.Values{}
	form.Add("campaign_code", campaign_code)
	form.Add("name", name)
	form.Add("phone", phone)
	form.Add("custom_data", customData)
	form.Add("instant", "true")

	req, err := http.NewRequest("POST", apiHost+"api/_partners/campaigns/subscribe-by-code", strings.NewReader(form.Encode()))
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return map[string]interface{}{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Api-Key", apiKey)
	resp, err := hc.Do(req)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return map[string]interface{}{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body Read Failed: %s", err)
		return map[string]interface{}{}, err
	}
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal result
	var rslt map[string]interface{}
	err = json.Unmarshal(body, &rslt)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return map[string]interface{}{}, err
	}

	return rslt, nil
}
