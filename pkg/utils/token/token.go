package token

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func Create(ttl time.Duration, payload interface{}, privateKeyBase64 string) (string, error) {
	privateKey, err := base64.StdEncoding.DecodeString(privateKeyBase64)

	if err != nil {
		return "", fmt.Errorf("Create: could not decode base64 encoded key: %w", err)
	}

	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		return "", fmt.Errorf("Create: parse rsa key from PEM: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaKey)

	if err != nil {
		return "", fmt.Errorf("Create: creating signed token: %w", err)
	}

	return token, nil
}

func Validate(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("Validate: could not decode base64 key: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return "", fmt.Errorf("Validate: could not parse rsa key from PEM: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Validate: unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("Validate: invalid token")
	}

	return claims["sub"], nil
}
