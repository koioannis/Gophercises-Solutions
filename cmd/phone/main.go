package main

import (
	"log"
	"os"
	"unicode"

	"github.com/joho/godotenv"
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
	normalize(phoneNumbers)

	ph.UpdateMany(phoneNumbers)
	ph.DeleteDuplicates()
}

func normalize(phoneNumbers []models.PhoneNumber) {
	for i := 0; i < len(phoneNumbers); i++ {
		phoneNumbers[i].Number = normalizeStr(phoneNumbers[i].Number)
	}
}

func normalizeStr(number string) string {
	var normalized []byte

	for _, c := range number {
		if unicode.IsDigit(c) {
			normalized = append(normalized, byte(c))
		}
	}
	return string(normalized)
}
