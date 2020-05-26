package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "/users/login",
		ReqBody:      `{"email": "email@gmail.com","password": "the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, err.Message, "Invalid rest client response when trying to login user")
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "/users/login",
		ReqBody:      `{"email": "email@gmail.com","password": "the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials","status": "404","error": "not_found"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, err.Message, "invalid error interface when trying to login user")
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "/users/login",
		ReqBody:      `{"email": "email@gmail.com","password": "the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials","status": 404,"error": "not_found"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, err.Message, "invalid login credentials")
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "/users/login",
		ReqBody:      `{"email": "email@gmail.com","password": 1234}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1234","first_name": "first-name","last-name": "last-name","email":"email@gmail.com"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, err.Message, "Error when unmarshalling the response")
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "/users/login",
		ReqBody:      `{"email": "email@gmail.com","password": 1234}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1234,"first_name": "first-name","last_name": "last-name","email":"email@gmail.com"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1234, user.Id)
	assert.EqualValues(t, "first-name", user.FirstName)
	assert.EqualValues(t, "last-name", user.LastName)
	assert.EqualValues(t, "email@gmail.com", user.Email)
}
