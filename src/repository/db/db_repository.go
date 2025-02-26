package db

import (
	"github.com/dralos/bookstore_oauth-api/src/domain/access_token"
	"github.com/dralos/bookstore_oauth-api/src/utils/errors"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
