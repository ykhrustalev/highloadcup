package highloadcup

import (
	"net/http"
	"time"
)

func Server() {
	usersRepo := NewUsersRepoImpl()
	usersRepo.Save(&User{
		Id:        1,
		Email:     "email@goo.com",
		FirstName: "first",
		LastName:  "last",
		Gender:    "f",
		BirthDate: time.Now(),
	})
	usersHandler := NewUsersHandler(usersRepo)


	locationsRepo := NewLocationsRepoImpl()
	locationsHandler := NewLocationsHandler(locationsRepo)

	visitsRepo := NewVisitsRepoImpl()
	visitsHandler := NewVisitsHandler(visitsRepo)

	router := NewRouter(usersHandler, locationsHandler, visitsHandler)

	http.HandleFunc(usersHandler.path, router.Handle)
	http.ListenAndServe(":8080", nil)
}
