package domains

import "time"

type (
	Jwt struct {
		Sub string
		Exp time.Time
	}
)
