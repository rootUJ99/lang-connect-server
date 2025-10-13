package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	repository "github.com/rootUJ99/lang-connect-server/internal/platform/postgres"

	"github.com/rootUJ99/lang-connect-server/internal/language"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("true"))	
}


func main() {
	slog.Default()
	err:= godotenv.Load()
	if (err != nil){
		slog.Error("Can not load .env file", "error", err)
	}
	slog.Info("Helloooooo lang server")
	r := chi.NewRouter()

	conn_string := os.Getenv("GOOSE_DBSTRING")

	slog.Info("Connected to", "db string", conn_string)
	
	conn, err := pgx.Connect(context.Background(), conn_string)

	if err != nil {
		slog.Error("Postgres connection failed", "error", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	query := repository.New(conn)
	languageRepository := language.NewLangRepo(query)
	languaageService := language.NewLanguageService(languageRepository)
	languageHandler := language.NewLanguageHandler(languaageService)

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
	    AllowedOrigins:   []string{"https://*", "http://*"},
	    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	    ExposedHeaders:   []string{"Link"},
	    AllowCredentials: false,
	    MaxAge:           300, 
	  }))

	r.Get("/", rootHandler)

	rw := chi.NewRouter() 
	
	rw.Post("/create", languageHandler.CreateLanguageHandler)
	rw.Put("/update", languageHandler.UpdateLanguageHandler)
	rw.Delete("/delete", languageHandler.DeleteLanguageHandler)
	rw.Get("/list", languageHandler.ListLanguageHandler)

	r.Mount("/language", rw)
	
	const PORT string = ":3141"
	slog.Info("Listening on", "PORT", PORT)	
	http.ListenAndServe(PORT, r)
}
