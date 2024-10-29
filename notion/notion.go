package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"orphanjamaahmgr/model"
	"os"
)

var (
	notionUrl    = os.Getenv("NOTION_URL")
	notionDbMjm  = os.Getenv("NOTION_DB_MJM")
	notionSecret = os.Getenv("NOTION_SECRET")
)

func GetListJamaahByLeadsCreates(leadsCreated string) (model.NotionRespBody, error) {
	hc := http.Client{}

	payload := []byte(`
		{
			"filter": {
				"or": [
					{
						"property" : "Leads Created",
						"date" : {
								"after": "` + leadsCreated + `"
						}
					}
				]
			},
			"sorts": [
				{
					"property": "Leads Created",
					"direction": "descending"
				}
			],
			"page_size": 50
		}
	`)

	req, err := http.NewRequest("POST", notionUrl+"/v1/databases/"+notionDbMjm+"/query", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return model.NotionRespBody{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+notionSecret)
	req.Header.Add("Notion-Version", "2022-06-28")

	resp, err := hc.Do(req)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return model.NotionRespBody{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body Read Failed: %s", err)
		return model.NotionRespBody{}, err
	}

	// Unmarshal result
	var notionRespBody model.NotionRespBody
	err = json.Unmarshal(body, &notionRespBody)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		fmt.Println("emang apa body raw nya : ", string(body))
		return model.NotionRespBody{}, err
	}

	return notionRespBody, nil
}
