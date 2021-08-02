package repositories

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"gorm.io/gorm"
)

const (

)

var (
	CodeRepository codeRepositoryInterface = &codeRepository{}
)

type codeRepository struct {
	DB *gorm.DB
}

type codeRepositoryInterface interface {
	FindCode(phone string , code , reason int) (*domains.Code, rest_errors.RestErr)
}

func NewCodeRepository(db *gorm.DB) *codeRepository {
	return &codeRepository{DB: db}
}

func (c *codeRepository) FindCode(phone string , code , reason int) (*domains.Code, rest_errors.RestErr) {
	return nil, nil
}
