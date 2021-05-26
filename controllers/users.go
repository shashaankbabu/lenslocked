package controllers

import (
	"lenslocked/views"
	"net/http"
)

// NewUsers is used to create a new users controller.
// This function will panic  if the templates are not parsed correctly.
// Should only be used during intial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
