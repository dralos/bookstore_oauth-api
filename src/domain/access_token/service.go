package access_token

import (
	"strings"

	"github.com/dralos/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenid string) (*AccessToken, *errors.RestErr) {
	accessTokenid = strings.TrimSpace(accessTokenid)
	if len(accessTokenid) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenid)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
