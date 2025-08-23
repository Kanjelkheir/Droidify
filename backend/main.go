package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	"github.com/kanjelkheir/droidify/internal/agent"
	"github.com/kanjelkheir/droidify/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	serverHits atomic.Int32
	db         *database.Queries
}

func (cfg *apiConfig) getData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type request struct {
			Model string `json:"model"`
		}

		data, err := agent.GetData(r)
		if err != nil {
			fmt.Printf("Error communicating to ai model:\n %v", err)
		}

		jsonResponse, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})
}

func (cfg *apiConfig) hitsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.serverHits.Add(1)
		handler.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) getMetrics() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseData := struct {
			Hits int `json:"hits"`
		}{
			Hits: int(cfg.serverHits.Load()),
		}

		response, err := json.Marshal(responseData)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})
}

func main() {

	godotenv.Load()

	db_url := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		fmt.Println("Failed to open db:", err)
		return
	}

	dbQueries := database.New(db)

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	config := apiConfig{
		db: dbQueries,
	}

	mux.Handle("POST /api/root", config.hitsMiddleware(config.getData()))
	mux.Handle("GET /api/admin/hits", config.getMetrics())

	log.Println("Started listening on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
