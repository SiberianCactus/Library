package main

import (
	authors "books/Authors"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type AuthorsServer struct {
	store *authors.Authors_books_Store
}

func NewAuthorServer() *AuthorsServer {
	store := authors.New()
	return &AuthorsServer{store: store}
}

func (auth *AuthorsServer) AuthorsHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/authors/" {
		if req.Method == http.MethodPost {
			auth.createAuthorHandler(w, req)
		} else if req.Method == http.MethodGet {
			auth.getAllAuthorsHandler(w, req)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /authors/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	} else {
		path := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /authors/<id> in authors handler", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Method == http.MethodDelete {
			(auth.deleteAuthorHandler(w, req, id))
		} else if req.Method == http.MethodGet {
			(auth.getAuthorHandler(w, req, id))
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /authors/<id>, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (auth *AuthorsServer) createAuthorHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling author create at %s\n", req.URL.Path)

	type RequestAuthor struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Book    string `json:"book"`
	}

	type ResponseId struct {
		Id int `json:"id"`
	}

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var ra RequestAuthor
	if err := dec.Decode(&ra); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := auth.store.CreateAuthor(ra.Name, ra.Surname, ra.Book)
	js, err := json.Marshal(ResponseId{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func (auth *AuthorsServer) getAllAuthorsHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all authors at %s\n", req.URL.Path)

	allTasks := auth.store.GetAllAuthors()
	js, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (auth *AuthorsServer) getAuthorHandler(w http.ResponseWriter, req *http.Request, id int) {
	log.Printf("handling get task at %s\n", req.URL.Path)

	author, err := auth.store.GetAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (auth *AuthorsServer) deleteAuthorHandler(w http.ResponseWriter, req *http.Request, id int) {
	log.Printf("handling delete author at %s\n", req.URL.Path)

	err := auth.store.DeleteAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func main() {
	mux := http.NewServeMux()
	server := NewAuthorServer()
	mux.HandleFunc("/authors/", server.AuthorsHandler)

	log.Fatal(http.ListenAndServe("localhost:8080"+os.Getenv("SERVERPORT"), mux))
}
