package phone

import (
	"unicode"

	"github.com/koioannis/gophercises-solutions/phone/models"
)

type Normalizer interface {
	Normalize(phoneNumbers []models.PhoneNumber)
}

func NewSimpleNormalizer() *SimpleNormalizer {
	return &SimpleNormalizer{}
}

type SimpleNormalizer struct{}

func (sn *SimpleNormalizer) Normalize(phoneNumbers []models.PhoneNumber) {
	for i := 0; i < len(phoneNumbers); i++ {
		phoneNumbers[i].Number = sn.normalizeStr(phoneNumbers[i].Number)
	}
}

func (sn *SimpleNormalizer) normalizeStr(number string) string {
	var normalized []byte

	for _, c := range number {
		if unicode.IsDigit(c) {
			normalized = append(normalized, byte(c))
		}
	}
	return string(normalized)
}
