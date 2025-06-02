package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func GetTokenSubject(jwt string) JWTSubject {
	var jwtElement = strings.Split(strings.Trim(jwt, "Bearer "), ".")

	var payload Payload

	payloadJson, _ := base64.RawURLEncoding.DecodeString(jwtElement[1])

	json.Unmarshal(payloadJson, &payload)

	return payload.Sub
}

func GenerateToken(user JWTSubject) string {
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

func GenerateRefreshToken(user JWTSubject) string {
	var secretKey = os.Getenv("SECRET_KEY")

	var header Header = Header{
		Alg: "sha256",
		Typ: "jwt",
	}

	headerJson, _ := json.Marshal(header)
	headerEncode := base64.RawURLEncoding.EncodeToString(headerJson)

	currentTime := GetCurrentTimeVN()

	// ⚠️ Expiry time for refresh token: 7 days
	tokenExp := currentTime.Add(7 * 24 * time.Hour).Format("02-01-2006 15:04:05")

	var payload Payload = Payload{
		Sub: user,
		Exp: tokenExp,
	}

	payloadJson, _ := json.Marshal(payload)
	payloadEncode := base64.RawURLEncoding.EncodeToString(payloadJson)

	signature := Signature{
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

func GetCurrentTimeVN() time.Time {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("⚠️ Lỗi khi load location:", err)
		return time.Now().UTC() // Fallback về UTC nếu load location lỗi
	}

	return time.Now().In(location)
}
