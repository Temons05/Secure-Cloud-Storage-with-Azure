package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"fmt"

	db "github.com/Ovenoboyo/basic_webserver/pkg/database"
	"github.com/Ovenoboyo/basic_webserver/pkg/handlers"
	"github.com/Ovenoboyo/basic_webserver/pkg/middleware"
	"github.com/Ovenoboyo/basic_webserver/pkg/storage"
	"github.com/joho/godotenv"
	"github.com/markbates/pkger"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env")
	}

	db.ConnectToDB()
	storage.InitializeStorage()

	pkger.Include("/static")

	r := mux.NewRouter()
	apiRouter := mux.NewRouter()
	apiRouterNegroni := middleware.GetJWTWrappedNegroni(apiRouter)

	r.PathPrefix("/api").Handler(apiRouterNegroni)
	http.Handle("/", r)

	handlers.HandleBlobs(apiRouter)
	handlers.HandleLogin(r)
	handlers.HandleStatic(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	})

	n := negroni.Classic()
	n.Use(c)
	apiRouterNegroni.Use(c)
	n.UseHandler(r)

	port := "8081"
	if os.Getenv("DEBUG") == "FALSE" {
		port = "80"
	}

	fmt.Println("here")
	n.Run(":" + port)
}
