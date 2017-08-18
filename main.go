package highloadcup

import (
	"net/http"
	"time"
)

func Server() {
	userRepo := NewUserRepoImpl()
	userRepo.Add(&User{
		Id: 1,
		Email: "x@y.z",
		FirstName: "first",
		LastName: "last",
		Gender: "m",
		BirthDate: time.Now(),
	})
	userHandler := NewUsersHandler(userRepo)

	http.HandleFunc(userHandler.Path, userHandler.Handle)
	http.ListenAndServe(":8080", nil)
}
