package controllers

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/dto"
	"net/http"
)

func NewProfileController() ProfileController {
	return ProfileController{}
}

type ProfileController struct {
}

func (c *ProfileController) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c.Update(w, r)
	} else {
		dto.ToJsonResponse(w, dto.UserResponseFromModel(*config.CurrentUser), http.StatusOK)
	}
}

func (c *ProfileController) Update(w http.ResponseWriter, r *http.Request) {
	userRequest := dto.UserRequest{}
	dto.FromRequest(r, &userRequest)
	(*config.CurrentUser).Name = userRequest.Name
	(*config.CurrentUser).Email = userRequest.Email
	config.DB.Save(config.CurrentUser)
	dto.ToStatusResponse(w, http.StatusOK)
}
