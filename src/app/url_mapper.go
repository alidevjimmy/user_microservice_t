package app

import (
	"fmt"
	"github.com/alidevjimmy/user_microservice_t/src/controllers/v1"
)

const (
	v1Prefix = "/v1/%s"
)

func urlMapper() {
	// v1
	e.GET(fmt.Sprintf(v1Prefix ,"register") , v1.Register)
}