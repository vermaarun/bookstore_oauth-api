package db

import (
	"github.com/vermaarun/bookstore_oauth-api/src/clients/cassandra"
	"github.com/vermaarun/bookstore_oauth-api/src/domain/access_token"
	"github.com/vermaarun/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	// TODO: implement
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("DB not implemented yet..!!")
}
