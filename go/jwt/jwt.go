package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/krishak-fiem/utils/go/env"
	"time"
)

const duration = 24 * time.Hour

type Payload struct {
	ID        uint64    `json:"id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p Payload) Valid() error {
	if p.ID > 0 {
		return errors.New("Invalid ID.")
	}

	if time.Now().After(p.ExpiredAt) {
		return errors.New("Token Expired.")
	}

	return nil
}

func newPayload(id uint64) (*Payload, error) {
	payload := &Payload{
		ID:        id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func CreateToken(id uint64) (string, error) {
	payload, err := newPayload(id)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(env.Get("JWT_SECRET")))
}

func VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token.")
		}
		return []byte(env.Get("JWT_SECRET")), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, errors.New("Invalid Token.")
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("Invalid Token.")
	}

	return payload, nil
}
