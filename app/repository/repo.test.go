package repository

import (
	"app/domain/model"
	"app/internal/db"
	"fmt"
)

func TestCampaignRepo() {

	repo := new(CampaignRepository)
	repo.Init(db.GetDB())

	campaign := model.Campaign{}

	fmt.Println(repo.Collection())

	err := repo.Create(&campaign, nil)

	if err != nil {

		fmt.Println(err)
	}
}
