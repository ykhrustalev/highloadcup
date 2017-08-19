package highloadcup

import (
	"net/http"
	"os"
	"fmt"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
