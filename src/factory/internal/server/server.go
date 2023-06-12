package server

import (
	"encoding/json"
	"fmt"
	"github.com/hramov/tbotfactory/internal/robot"
	"io"
	"log"
	"net/http"
)

type Server interface {
	Start() error
}

type server struct {
	port int
}

func New(port int) Server {
	return &server{port: port}
}

func (s *server) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong http method", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("cannot send response: %v", err)
	}
}

func (s *server) new(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusBadRequest)
		return
	}

	type bot struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var b bot

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		if err == io.EOF {
			http.Error(w, "no body", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if b.Name == "" || b.Title == "" || b.Description == "" {
		http.Error(w, "no required fields found", http.StatusBadRequest)
		return
	}

	creator := robot.New()
	token, err := creator.Create(b.Name, b.Title, b.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf("{ token: %s}", token)))
	if err != nil {
		log.Printf("cannot send response: %v", err)
	}
}

func (s *server) Start() error {

	http.HandleFunc("/health", s.health)
	http.HandleFunc("/new", s.new)

	port := fmt.Sprintf(":%d", s.port)

	log.Printf("starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}
