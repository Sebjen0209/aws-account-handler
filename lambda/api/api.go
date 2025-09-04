package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("There are either missing a username or a password")
	}

	userExist, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("There was an error checking if the user exists: %w", err)
	}

	if userExist {
		return fmt.Errorf("The user does already exist with that username")
	}

	//we know that a user does not exist
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("Something happened registering the user, error: %w", err)
	}
	return nil
}
