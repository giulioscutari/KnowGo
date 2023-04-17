package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Run() any {
	router := http.NewServeMux()
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
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/document/"))
	doc := NewDocument(id, "document title", "lorem. Ipsum. Wow. Test")
	WriteJSON(w, http.StatusOK, doc)
}

func (s *Server) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, "DeleteDocument")
}

func (s *Server) PostDocument(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, "PostDocument")
}
