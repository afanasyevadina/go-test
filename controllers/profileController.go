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
	dto.ToJsonResponse(w, dto.UserResponseFromModel(*config.CurrentUser), http.StatusOK)
}
