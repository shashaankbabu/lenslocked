package controllers

import (
	"fmt"
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

// New() is used to render a form for account creation.
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create() is used to process a signup form.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is temp response.")
}
