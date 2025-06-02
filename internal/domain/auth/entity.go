package auth

type AuthLogin struct {
	Email    string
	Password string
	Device   string
}

type JWTSubject struct {
	Id      uint
	Device  string
	Version uint
}

type Header struct {
	Alg string
	Typ string
}

type Payload struct {
	Sub JWTSubject
	Exp string
}

type Signature struct {
	HeaderEncode  string
	PayloadEncode string
}
