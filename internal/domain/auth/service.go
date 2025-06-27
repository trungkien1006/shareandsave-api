package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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

func (s *AuthService) GetTokenSubject(jwt string) JWTSubject {
	var jwtElement = strings.Split(strings.Trim(jwt, "Bearer "), ".")

	var payload Payload

	payloadJson, _ := base64.RawURLEncoding.DecodeString(jwtElement[1])

	json.Unmarshal(payloadJson, &payload)

	return payload.Sub
}

func (s *AuthService) GenerateToken(user JWTSubject) string {
	var secretKey = os.Getenv("SECRET_KEY")

	var header Header = Header{
		Alg: "sha256",
		Typ: "jwt",
	}

	headerJson, _ := json.Marshal(header)

	var headerEncode = base64.RawURLEncoding.EncodeToString(headerJson)

	currentTime := GetCurrentTimeVN()

	tokenExp := currentTime.Add(time.Hour * 4).Format("02-01-2006 15:04:05")

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

func (s *AuthService) GenerateRefreshToken(user JWTSubject) string {
	var secretKey = os.Getenv("SECRET_KEY")

	var header Header = Header{
		Alg: "sha256",
		Typ: "jwt",
	}

	headerJson, _ := json.Marshal(header)
	headerEncode := base64.RawURLEncoding.EncodeToString(headerJson)

	currentTime := GetCurrentTimeVN()

	// ⚠️ Expiry time for refresh token: 30 days
	tokenExp := currentTime.Add(30 * 24 * time.Hour).Format("02-01-2006 15:04:05")

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

func (s *AuthService) GenerateEmailToken(email string) string {
	base := email + time.Now().Format("20060102150405")
	hash := sha256.Sum256([]byte(base))
	return hex.EncodeToString(hash[:])[:6] // Lấy 6 ký tự đầu của hash
}

func GetCurrentTimeVN() time.Time {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("⚠️ Lỗi khi load location:", err)
		return time.Now().UTC() // Fallback về UTC nếu load location lỗi
	}

	return time.Now().In(location)
}

func (s *AuthService) BuildOTPEmailContent(otp, sub string) (subject, body string) {
	subject = sub

	body = fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<style>
				body { font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; }
				.container { max-width: 600px; margin: 0 auto; background-color: #fff; border-radius: 10px; padding: 30px; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
				.title { color: #333; font-size: 24px; font-weight: bold; margin-bottom: 20px; }
				.otp { font-size: 32px; font-weight: bold; color: #007bff; letter-spacing: 4px; margin: 20px 0; }
				.footer { font-size: 14px; color: #888; margin-top: 30px; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="title">Xin chào,</div>
				<p>Bạn vừa yêu cầu xác thực tài khoản. Đây là mã OTP của bạn:</p>
				<div class="otp">%s</div>
				<p>Mã này sẽ hết hạn sau 5 phút. Vui lòng không chia sẻ mã này với bất kỳ ai.</p>
				<div class="footer">Cảm ơn bạn đã sử dụng dịch vụ của chúng tôi.</div>
			</div>
		</body>
		</html>
	`, otp)

	return
}

// GenerateVerificationToken tạo token duy nhất từ email
func (s *AuthService) GenerateVerificationToken(email string) (string, error) {
	// Tạo random salt 16 bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Kết hợp email + timestamp + salt
	data := fmt.Sprintf("%s|%d|%s", email, time.Now().UnixNano(), base64.URLEncoding.EncodeToString(salt))

	// Băm bằng SHA-256
	hash := sha256.Sum256([]byte(data))

	// Trả về token dạng base64
	token := base64.URLEncoding.EncodeToString(hash[:])

	return token, nil
}
