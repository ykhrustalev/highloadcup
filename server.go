package highloadcup

import (
	"net/http"
)

func Server() {
	userRepo := NewUsersRepoImpl()
	usersHandler := NewUsersHandler(userRepo)

	router := NewRouter(usersHandler)

	http.HandleFunc(usersHandler.path, router.Handle)
	http.ListenAndServe(":8080", nil)
}
