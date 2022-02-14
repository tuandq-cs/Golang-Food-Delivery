package jwt

import (
	"Golang_Edu/component/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtTokenProvider struct {
	secret string
}

func NewJWTTokenProvider(secret string) *jwtTokenProvider {
	return &jwtTokenProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtTokenProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
		},
	})
	accessToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	return &tokenprovider.Token{
		Token:   accessToken,
		Expiry:  expiry,
		Created: now,
	}, nil

}

func (j *jwtTokenProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	t, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil || !t.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}
	claims, ok := t.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &claims.Payload, nil
}
