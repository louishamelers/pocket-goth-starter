package auth

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func RegisterUser(e *core.RequestEvent, email string, password string, repeatPassword string) error {
	user, _ := e.App.FindAuthRecordByEmail("users", email)
	if user != nil {
		return fmt.Errorf("A user with that email already exists")
	}

	// TODO: move this to validation
	if repeatPassword != password {
		return fmt.Errorf("The entered passwords do not match")
	}

	userCollection, err := e.App.FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	newUser := core.NewRecord(userCollection)
	newUser.SetPassword(password)
	newUser.SetEmail(email)

	if err := e.App.Save(newUser); err != nil {
		return err
	}

	return SetAuthToken(e, newUser)
}

func LoginUser(e *core.RequestEvent, email string, password string) error {
	user, err := e.App.FindAuthRecordByEmail("users", email)
	if err != nil {
		return fmt.Errorf("Login failed")
	}

	valid := user.ValidatePassword(password)

	if !valid {
		return fmt.Errorf("Login failed")
	}

	return SetAuthToken(e, user)
}

func LogoutUser(e *core.RequestEvent) {
	RemoveAuthToken(e)
}
