package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
	"os"
)

func main() {

	database.InitDB()
	db := database.GetDB()
	modelCreated := prepareModel()

	result := db.Create(&modelCreated)
	if result.Error != nil {
		fmt.Println("Error in insert data to DB " + result.Error.Error())
		os.Exit(1)
	}


	fmt.Printf("Row count added %v " , result.RowsAffected)
}

func prepareModel() *model.Device  {
	device := &model.Device{}
	err := faker.FakeData(device)
	if err != nil {
		fmt.Println("Error in faker data " + err.Error())
		os.Exit(1)
	}

	return  device
}






