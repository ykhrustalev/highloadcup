package highloadcup

import (
	"net/http"
	"os"
	"fmt"
	"log"
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
		path = "/tmp/data/data.zip"
	}

	loader := NewLoader(usersRepo, locationsRepo, visitsRepo)
	err := loader.Load(path)
	if err != nil {
		log.Fatalf("failed to load data, %v", err)
	}

	router := NewRouter(usersHandler, locationsHandler, visitsHandler)

	http.HandleFunc("/", router.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("listen on %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
