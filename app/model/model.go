package model

import (
	"app/app/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type IModel interface {
	Store()
	Insert()
	Update()
	Delete()
}

type Model struct {
	db_client *mongo.Client
}

func (this *Model) Init() {

	this.initDBClient()
}

func (this *Model) Establish() {

	this.Init()
}

func (this *Model) initDBClient() {

	if this.db_client != nil {

		return
	}

	client, _ := db.GetClient()

	this.db_client = client
}
