package util

import "golang.org/x/crypto/bcrypt"

const (
	cost = 10
)

type UserPasswordEncryptor interface {
	Encrypt(text string) (string, error)
}

type EncryptionService struct{}

func (e *EncryptionService) Encrypt(text string) (string, error) {

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(text), cost)
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}
