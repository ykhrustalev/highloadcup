package highloadcup

import (
	"net/http"
	"time"
)

func Server() {
	userRepo := NewUsersRepoImpl()
	userRepo.Save(&User{
		Id:        1,
		Email:     "x@y.z",
		FirstName: "first",
		LastName:  "last",
		Gender:    "m",
		BirthDate: time.Now(),
	})
	usersHandler := NewUsersHandler(userRepo)

	router := NewRouter(usersHandler)

	http.HandleFunc(usersHandler.path, router.Handle)
	http.ListenAndServe(":8080", nil)
}
