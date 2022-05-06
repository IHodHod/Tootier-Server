package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
	"math/rand"
)

const NumberOfInsertModel =5

func main() {
	database.InitDB()
	db := database.GetDB()

	tx := db.Begin()
	for i := 0; i < NumberOfInsertModel; i++ {
		user := model.User{
			UserName:     faker.Username(),
			Password:     faker.Password(),
			Name:         faker.Name(),
			Email:        faker.Email(),
			PhoneNumber:  faker.Phonenumber(),
			RegisterTime: faker.Timestamp(),
			UserGeo: model.UserGeo{
				LatitudeX:  fmt.Sprintf("%f", faker.Latitude()),
				LongitudeY: fmt.Sprintf("%f", faker.Longitude()),
			},
			ProfileID: uint64(i + 1),
		}
		tx.Create(&user)

		device := model.Device{
			UserID:             uint64(rand.Intn(NumberOfInsertModel / 2)),
			LastTimeVisit:      faker.Timestamp(),
			RegisterTimeDevice: faker.Timestamp(),
			Token:              faker.Jwt(),
			Hardware: model.Hardware{
				DeviceName: faker.UUIDDigit(),
				MacAddress: faker.MacAddress(),
				OS:         faker.UUIDDigit(),
				IP:         faker.IPv4(),
				DeviceUUID: faker.UUIDDigit(),
			},
		}

		tx.Create(&user)
		tx.Create(&device)
	}
	tx.Commit()
}
