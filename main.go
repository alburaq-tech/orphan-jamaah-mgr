package main

import (
	"fmt"
	"orphanjamaahmgr/langgeng"
	"orphanjamaahmgr/notion"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("oke")
	args := os.Args

	// fmt.Println("agrgs: ", argsWithProg[1])

	if len(args) <= 2 {
		fmt.Println("dont forget the argument")
		return
	}

	leadsCreatedDay := args[1] + "T00:00:00+07:00"
	windowEligibleRaw := args[2]
	windowEligible, err := strconv.Atoi(windowEligibleRaw)
	if err != nil {
		fmt.Println("error on parsing windowEligibleRaw : ", err)
		return
	}

	jamaahs, err := notion.GetListJamaahByLeadsCreates(leadsCreatedDay)
	if err != nil {
		fmt.Println("error on retrieving list orphan jamaah : ", err)
		return
	}

	listOrphanCount := 0
	listOrphanEligibleCount := 0
	listJamaah := 0
	for _, v := range jamaahs.Results {
		actualLeadsCreated := v.Properties.LeadsCreated.Date.Start

		if len(v.Properties.AccountOwner.People) == 0 {
			listOrphanCount++
			diff := time.Now().Sub(actualLeadsCreated)

			if diff < time.Duration(windowEligible)*time.Minute {
				listOrphanEligibleCount++
				// Print hours, minutes and seconds
				fmt.Printf("eligible for remind %s %.1f min\n", v.Properties.Name.Title[0].Text.Content, diff.Minutes())

				respLanggeng, err := langgeng.CallLanggengSubscribeCampaign("masteralburaq-reminder-join-group", v.Properties.Name.Title[0].Text.Content, v.Properties.NoWA08Xxxxx.PhoneNumber, "[]")
				time.Sleep(5 * time.Second)
				if err != nil {
					fmt.Println("error on calling langgeng subscribe campaign ", err)
				} else {
					fmt.Println("WHAT IS THE RESULT FROM LANGGENG : ", respLanggeng)
				}
			}
		}
		listJamaah++
	}

	fmt.Println("how many jamaah today : ", listJamaah)
	fmt.Println("how many orphan today : ", listOrphanCount)
	fmt.Println("how many orphan elgigible for remind: ", listOrphanEligibleCount)
	// fmt.Println("sample orphan : ", string(sampleOrphan))

}
