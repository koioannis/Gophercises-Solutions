package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/koioannis/gophercises-solutions/phone"
	"github.com/koioannis/gophercises-solutions/phone/models"
)

func main() {
	err := godotenv.Load("phone/.env")
	if err != nil {
		log.Fatal(err)
	}

	ph, err := models.NewPhoneNumbers(os.Getenv("CONN_STR"))
	if err != nil {
		log.Fatal(err)
	}
	ph.AutoMigrate()

	phoneNumbers, err := ph.GetAll()
	if err != nil {
		log.Fatal(phoneNumbers)
	}

	sn := phone.NewSimpleNormalizer()
	sn.Normalize(phoneNumbers)

	ph.UpdateMany(phoneNumbers)
	ph.DeleteDuplicates()
}
