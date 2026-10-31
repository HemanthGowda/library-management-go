package password

import "golang.org/x/crypto/bcrypt"

type Encoder interface {
	Encode(raw string) (string, error)
	Matches(raw, encoded string) bool
}

type Bcrypt struct{}

func (Bcrypt) Encode(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (Bcrypt) Matches(raw, encoded string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encoded), []byte(raw)) == nil
}
