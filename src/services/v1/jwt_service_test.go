package services

import (
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	Secret = "secret"
)

func TestFailToSignString(t *testing.T) {
	c := jwt.MapClaims{
		"sub": "ok",
		"exp": true,
	}
	token, err := JwtService.GenerateJwtToken(c)
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestGenerateJwtTokenSuccessfully(t *testing.T) {
	c := jwt.MapClaims{
		"sub": "1",
		"exp": time.Now().Add(time.Hour),
	}

	token, err := JwtService.GenerateJwtToken(c)
	assert.NotEqual(t, "", token)
	assert.Nil(t, err)
}

func TestVerifyJwtTokenFailToParseToken(t *testing.T) {
	token := "iuu2hwjelkqme,mne,dmndw"
	r, ok, err := JwtService.VerifyJwtToken(token)
	assert.Nil(t, r)
	assert.NotNil(t, err)
	assert.Equal(t, false, ok)
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestVerifyJwtTokenInvalidHashMethod(t *testing.T) {
	// SigningMethodEdDSA isn't HMAC
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"sub": "1",
		"exp": time.Now().Add(time.Hour),
	})
	tokenString, err := token.SignedString(Secret)
	assert.Nil(t, err)

	j, ok, err1 := JwtService.VerifyJwtToken(tokenString)

	assert.Nil(t, j)
	assert.NotNil(t, err1)
	assert.Equal(t, false, ok)
	assert.Equal(t, errors.InternalServerErrorMessage, err1.Message())
	assert.Equal(t, http.StatusInternalServerError, err1.Status())
}

func TestVerifyJwtTokenInvalidClaim(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour),
	})
	tokenString, err := token.SignedString(Secret)
	assert.Nil(t, err)

	j, ok, err1 := JwtService.VerifyJwtToken(tokenString)

	assert.Nil(t, j)
	assert.NotNil(t, err1)
	assert.Equal(t, false, ok)
	assert.Equal(t, errors.InternalServerErrorMessage, err1.Message())
	assert.Equal(t, http.StatusInternalServerError, err1.Status())
}

func TestVerifyJwtTokenUnableToUnMarshal(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"sub": "1",
	})
	tokenString, err := token.SignedString(Secret)
	assert.Nil(t, err)

	j, ok, err1 := JwtService.VerifyJwtToken(tokenString)

	assert.Nil(t, j)
	assert.NotNil(t, err1)
	assert.Equal(t, false, ok)
	assert.Equal(t, errors.InternalServerErrorMessage, err1.Message())
	assert.Equal(t, http.StatusInternalServerError, err1.Status())
}
