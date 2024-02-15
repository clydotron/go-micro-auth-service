package main

import (
	"errors"
	"fmt"
	"net/http"

	helpers "github.com/clydotron/toolbox/helpers"
)

func (a *App) Authenticate(w http.ResponseWriter, r *http.Request) {
	// TODO put these someplace common (but not in common)
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := helpers.ReadJSON(w, r, &requestPayload); err != nil {
		_ = helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := a.UserRepo.GetByEmail(requestPayload.Email)
	if err != nil {
		_ = helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		_ = helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := helpers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
	}
	_ = helpers.WriteJSON(w, http.StatusAccepted, payload)
}
