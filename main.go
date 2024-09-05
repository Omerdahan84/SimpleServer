package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main(){
	// load the enviorment variable port number 
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT not found")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: [] string {"https://*", "http://*"},
		AllowedMethods: [] string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: [] string{"*"},
		ExposedHeaders: [] string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	// creating new router, it will handle the /ready path
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // we want this to be used just with get method
	v1Router.Get("/err", handlerErr)
	
	// now we mount the original router on v1Router
	// so the full path is /v1/ready
	router.Mount("/v1",v1Router)
	// Set up the HTTP server using the router.
	// The server will listen on the port defined by the PORT environment variable.
	srv := &http.Server{
		Handler: router,            // Attach the router to handle incoming HTTP requests.
		Addr:    ":" + portString,  // Define the address and port to listen on.
	}

	log.Printf("Server is starting on port %v", portString)
	// Start the server and begin listening for requests.
	// If an error occurs, log it and terminate the program.
	err := srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Port: ", portString)
}