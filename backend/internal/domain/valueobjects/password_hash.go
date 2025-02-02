package valueobjects

import "golang.org/x/crypto/bcrypt"

type PasswordHash struct {
	value string
}

var encryptionCost = 10

func (p *PasswordHash) Encrypt(plainText string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), encryptionCost)

	if err != nil {
		panic(err)
	}

	p.value = string(hash)
}

func (p *PasswordHash) GetValue() string {
	return p.value
}

func (p *PasswordHash) Assign(value string) {
	p.value = value
}

func (p *PasswordHash) Validate(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(plainText))

	return err == nil
}
