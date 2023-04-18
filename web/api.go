package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	listenAddr string
	store      Storage
}

func NewServer(listenAddr string, store Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Run() any {
	router := http.NewServeMux()
	router.Handle("/documents/", http.HandlerFunc(s.HandleGetDocuments))
	router.Handle("/document/", http.HandlerFunc(s.HandleDocument))

	return (http.ListenAndServe(s.listenAddr, router))
}

func WriteJSON(w http.ResponseWriter, status int, v any) any {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) HandleDocument(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.GetDocument(w, r)
	case "POST":
		s.PostDocument(w, r)
	case "DELETE":
		s.DeleteDocument(w, r)
	}
}

func (s *Server) GetDocument(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimPrefix(r.URL.Path, "/document/")
	doc := NewDocument(title, "lorem. Ipsum. Wow. Test")
	WriteJSON(w, http.StatusOK, doc)
}

func (s *Server) HandleGetDocuments(w http.ResponseWriter, r *http.Request) {
	docs, err := s.store.GetDocuments()
	fmt.Println(docs)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	} else {
		WriteJSON(w, http.StatusOK, docs)
	}
}
func (s *Server) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, "DeleteDocument")
}

func (s *Server) PostDocument(w http.ResponseWriter, r *http.Request) {
	CreateDocReq := CreateDocumentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&CreateDocReq); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
	}

	document := NewDocument(CreateDocReq.Title, CreateDocReq.Body)

	if err := s.store.CreateDocument(document); err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)

	}

	WriteJSON(w, http.StatusOK, document)
}
