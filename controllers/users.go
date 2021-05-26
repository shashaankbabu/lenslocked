package controllers

import (
	"fmt"
	"lenslocked/views"
	"net/http"

	"github.com/gorilla/schema"
)

type Users struct {
	NewView *views.View
}

// NewUsers is used to create a new users controller.
// This function will panic  if the templates are not parsed correctly.
// Should only be used during intial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// New() is used to render a form for account creation.
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create() is used to process a signup form.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
