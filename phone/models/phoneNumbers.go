package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var numbers []string = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
	"(123) 456-7893",
	"123456-7894",
	"(123) 456-7894",
	"1234567894",
}

type PhoneNumber struct {
	gorm.Model
	Number string
}

func NewPhoneNumbers(dsn string) (*PhoneNumbers, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PhoneNumbers{
		db: db,
	}, nil
}

type PhoneNumbers struct {
	db *gorm.DB
}

func (pn *PhoneNumbers) GetAll() ([]PhoneNumber, error) {
	var phoneNumbers []PhoneNumber

	err := pn.db.Find(&phoneNumbers).Error
	if err != nil {
		return nil, err
	}

	return phoneNumbers, nil
}

func (pn *PhoneNumbers) UpdateMany(phoneNumbers []PhoneNumber) error {
	for _, phoneNumber := range phoneNumbers {
		err := pn.db.Save(&phoneNumber).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (pn *PhoneNumbers) DeleteDuplicates() error {
	return pn.db.Exec(`
	DELETE FROM
	phone_numbers a
	USING phone_numbers b
	WHERE
	a.id < b.id AND a.number = b.number`).Error
}

func (pn *PhoneNumbers) AutoMigrate() error {
	err := pn.db.AutoMigrate(&PhoneNumber{})
	if err != nil {
		return err
	}

	if pn.db.Find(&PhoneNumber{}).RowsAffected == 0 {
		return pn.createMockData()
	}

	return nil
}

func (pn *PhoneNumbers) createMockData() error {
	var mocks []PhoneNumber

	for _, numberStr := range numbers {
		mocks = append(mocks, PhoneNumber{
			Number: numberStr,
		})
	}

	for _, number := range mocks {
		err := pn.db.Create(&number).Error
		if err != nil {
			return err
		}
	}

	return nil
}
