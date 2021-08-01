package domains

import "time"

type (
	Jwt struct {
		Header  JwtHeader
		Payload JwtPayload
	}

	JwtHeader struct {
		Alg string
		Typ string
	}

	JwtPayload struct {
		Sub string
		Exp time.Time
	}
)
