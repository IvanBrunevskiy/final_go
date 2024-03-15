package main

import (
	"context"
	"encoding/json"
	"final_project/pkg/hasher"
	"google.golang.org/grpc"
	"log"
	http "net/http"
)

type server struct {
	hashClient hasher.HashingServiceClient
}

func NewServer(hashAddr string) (*server, error) {
	conn, err := grpc.Dial(hashAddr, grpc.WithInsecure()) // In production, use secure connection
	if err != nil {
		return nil, err
	}
	hashClient := hasher.NewHashingServiceClient(conn)
	return &server{hashClient: hashClient}, nil
}

func (s *server) checkHashHandler(w http.ResponseWriter, r *http.Request) {
	// Предположим, что 'hash' передается как параметр запроса
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "Query parameter 'hash' is missing", http.StatusBadRequest)
		return
	}

	// Вызываем метод gRPC сервиса
	resp, err := s.hashClient.CheckHash(context.Background(), &hasher.HashRequest{Payload: hash})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту
	json.NewEncoder(w).Encode(resp)
}

func (s *server) getHashHandler(w http.ResponseWriter, r *http.Request) {
	// Предположим, что 'payload' передается как параметр запроса
	payload := r.URL.Query().Get("payload")
	if payload == "" {
		http.Error(w, "Query parameter 'payload' is missing", http.StatusBadRequest)
		return
	}

	resp, err := s.hashClient.GetHash(context.Background(), &hasher.HashRequest{Payload: payload})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *server) createHashHandler(w http.ResponseWriter, r *http.Request) {
	// Предположим, что 'payload' передается в теле запроса
	var request hasher.HashRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.hashClient.CreateHash(context.Background(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201 Created для нового ресурса
	json.NewEncoder(w).Encode(resp)
}

func main() {
	srv, err := NewServer("localhost:50051") // Replace with actual Hashing Service address
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	http.HandleFunc("/checkhash", srv.checkHashHandler)
	http.HandleFunc("/gethash", srv.getHashHandler)
	http.HandleFunc("/createhash", srv.createHashHandler)

	log.Println("Gateway Service running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
