package services

import (
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestFailToSignString(t *testing.T) {
	c := jwt.MapClaims{
		"sub":"ok",
		"exp":true,
	}
	token , err := JwtService.GenerateJwtToken(c)
	assert.Equal(t , "",token)
	assert.NotNil(t, err)
	assert.Equal(t , errors.InternalServerErrorMessage, err.Message())
	assert.Equal(t ,http.StatusInternalServerError, err.Status())
}

func TestGenerateJwtTokenSuccessfully(t *testing.T) {
	c := jwt.MapClaims{
		"sub":"1",
		"exp": time.Now().Add(time.Hour),
	}
	token , err := JwtService.GenerateJwtToken(c)
	assert.Equal(t , "",token)
	assert.NotNil(t, err)
	assert.Equal(t , errors.InternalServerErrorMessage, err.Message())
	assert.Equal(t ,http.StatusInternalServerError, err.Status())
}
