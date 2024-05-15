package repository

import (
	"app/app/db"
	"app/app/model"
	"fmt"
)

func TestCampaignRepo() {

	repo := new(CampaignRepository)
	repo.Init(db.GetDB())

	campaign := model.Campaign{}

	fmt.Println(repo.Collection())

	err := repo.Create(&campaign)

	if err != nil {

		fmt.Println(err)
	}
}
