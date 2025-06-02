package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type UserJWTSubject struct {
	Id      uint
	Device  string
	Version uint
}

type Header struct {
	Alg string
	Typ string
}

type Payload struct {
	Sub UserJWTSubject
	Exp string
}

type Signature struct {
	HeaderEncode  string
	PayloadEncode string
}

func GetTokenSubject(jwt string) UserJWTSubject {
	var jwtElement = strings.Split(strings.Trim(jwt, "Bearer "), ".")

	var payload Payload

	payloadJson, _ := base64.RawURLEncoding.DecodeString(jwtElement[1])

	json.Unmarshal(payloadJson, &payload)

	return payload.Sub
}

func GenerateToken(user UserJWTSubject) string {
	var secretKey = os.Getenv("SECRET_KEY")

	var header Header = Header{
		Alg: "sha256",
		Typ: "jwt",
	}

	headerJson, _ := json.Marshal(header)

	var headerEncode = base64.RawURLEncoding.EncodeToString(headerJson)

	currentTime := GetCurrentTimeVN()

	tokenExp := currentTime.Add(time.Hour * 1).Format("02-01-2006 15:04:05")

	var payload Payload = Payload{
		Sub: user,
		Exp: tokenExp,
	}

	payloadJson, _ := json.Marshal(payload)

	var payloadEncode = base64.RawURLEncoding.EncodeToString(payloadJson)

	var signature Signature = Signature{
		HeaderEncode:  headerEncode,
		PayloadEncode: payloadEncode,
	}

	signatureJson, _ := json.Marshal(signature)

	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write(signatureJson)

	signatureHmac := h.Sum(nil)

	signatureEncode := base64.RawURLEncoding.EncodeToString(signatureHmac)

	token := fmt.Sprintf("%s.%s.%s", headerEncode, payloadEncode, signatureEncode)

	return token
}

func CheckJWT(jwt string) error {
	var secretKey = os.Getenv("SECRET_KEY")

	var jwtElement = strings.Split(strings.Replace(jwt, "Bearer", "", -1), ".")

	if jwt == "" || len(jwtElement) != 3 {
		return errors.New("token not found")
	}

	// Start check valid token
	var signature Signature = Signature{
		HeaderEncode:  strings.TrimSpace(jwtElement[0]),
		PayloadEncode: strings.TrimSpace(jwtElement[1]),
	}

	signatureJson, _ := json.Marshal(signature)

	signatureString := strings.Replace(string(signatureJson), "\\", "", -1)

	h := hmac.New(sha256.New, []byte(secretKey))

	h.Write([]byte(signatureString))

	signatureHmac := h.Sum(nil)

	signatureEncode := base64.RawURLEncoding.EncodeToString(signatureHmac)

	if check := signatureEncode == jwtElement[2]; !check {
		return errors.New("token not valid")
	}
	// End check valid token

	// Start check exp token
	payloadJson, _ := base64.RawURLEncoding.DecodeString(jwtElement[1])

	var payload Payload

	json.Unmarshal(payloadJson, &payload)

	exp, _ := time.Parse("02-01-2006 15:04:05", payload.Exp)

	currentTime := GetCurrentTimeVN()

	if checkTime := currentTime.Before(exp); !checkTime {
		return errors.New("token expired")
	}
	// End check exp token

	return nil
}
