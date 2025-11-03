package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type JWTPayload struct {
	Username string `json:"username"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"type"`
}

func GenerateToken(username string) (string, error) {
	header := JWTHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	now := time.Now()

	payload := JWTPayload{
		Username: username,
		Iat:      now.Unix(),
		Exp:      now.Add(1 * time.Hour).Unix(),
	}

	headerJSON, _ := json.Marshal(header)
	payloadJSON, _ := json.Marshal(payload)

	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadJSON)

	signature := sign(headerEncoded+"."+payloadEncoded, "my_secret")

	token := headerEncoded + "." + payloadEncoded + "." + signature
	return token, nil

}

func VerifyToken(token string) (bool, string) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, errors.New("Invalid token").Error()
	}

	headerEncoded := parts[0]
	payloadEncoded := parts[1]
	signature := parts[2]

	headerJSON, err := base64.RawURLEncoding.DecodeString(headerEncoded)
	if err != nil {
		return false, errors.New("Invalid token").Error()
	}
	var header JWTHeader
	err = json.Unmarshal(headerJSON, &header)
	if err != nil {
		return false, errors.New("Invalid token").Error()
	}
	if header.Alg != "HS256" {
		return false, errors.New("Invalid token").Error()
	}
	if header.Typ != "JWT" {
		return false, errors.New("Invalid token").Error()
	}
	payloadJSON, err := base64.RawURLEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return false, errors.New("Invalid token").Error()
	}
	var payload JWTPayload
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return false, errors.New("Invalid token").Error()
	}
	if payload.Exp < time.Now().Unix() {
		return false, errors.New("Token expired").Error()
	}
	signatureToCheck := sign(headerEncoded+"."+payloadEncoded, "my_secret")
	if signature != signatureToCheck {
		return false, errors.New("Invalid token").Error()
	}
	return true, ""

}

func sign(data string, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
