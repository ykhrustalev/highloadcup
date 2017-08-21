package highloadcup

import (
	"fmt"
	"github.com/ykhrustalev/highloadcup/data_loader"
	"github.com/ykhrustalev/highloadcup/repos"
	"log"
	"net/http"
	"os"
	"github.com/ykhrustalev/highloadcup/handlers"
	"github.com/ykhrustalev/highloadcup/handlers/crud"
	"time"
)

func Server() {
	t0 := time.Now()

	repo := repos.NewRepo()

	path := os.Getenv("DATA_PATH")
	if path == "" {
		path = "/tmp/data/data.zip"
	}

	loader := data_loader.NewLoader(repo)
	err := loader.Load(path)
	if err != nil {
		log.Fatalf("failed to load data, %v", err)
	}

	router := NewRouter(
		crud.NewHandler(repo),
		handlers.NewListVisitsHandler(repo),
		handlers.NewLocationsAvgHandler(repo),
	)

	http.HandleFunc("/", router.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	t1 := time.Now()
	fmt.Printf("booted in %d seconds\n", t1.Unix() - t0.Unix())

	fmt.Printf("listen on %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
