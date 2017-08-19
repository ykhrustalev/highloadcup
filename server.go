package highloadcup

import (
	"net/http"
	"os"
)

func Server() {
	usersRepo := NewUsersRepoImpl()
	usersHandler := NewUsersHandler(usersRepo)

	locationsRepo := NewLocationsRepoImpl()
	locationsHandler := NewLocationsHandler(locationsRepo)

	visitsRepo := NewVisitsRepoImpl()
	visitsHandler := NewVisitsHandler(visitsRepo)

	path := os.Getenv("DATA_PATH")
	if path == "" {
		path = "/tmp/data/data.zip "
	}

	loader := NewLoader(usersRepo, locationsRepo, visitsRepo)
	loader.Load(path)

	router := NewRouter(usersHandler, locationsHandler, visitsHandler)

	http.HandleFunc(usersHandler.path, router.Handle)
	http.ListenAndServe(":8080", nil)
}
