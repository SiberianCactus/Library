package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	flags "github.com/jessevdk/go-flags"

	"books/models"
	"books/restapi"
	"books/restapi/operations"
	"books/restapi/operations/authors"
)

type Handler struct {
	sync.RWMutex
	authors models.Authors
	nextId  int64
}

func (auth *Handler) Handle(authors.GetAllAuthorsParams) middleware.Responder {
	auth.RLock()
	defer auth.RUnlock()

	allAuthors := make(models.Authors, 0, len(auth.authors))
	for _, task := range auth.authors {
		allAuthors = append(allAuthors, task)
	}

	// return &authors.GetAllAuthorsDefault{
	// 	Payload: &models.Error{
	// 		Code:    swag.Int64(http.StatusBadRequest),
	// 		Message: swag.String("Bad Request"),
	// 	},
	// }

	return &authors.GetAllAuthorsOK{
		Payload: allAuthors,
	}
}

func (auth *Handler) AddAuthorHandler(request authors.AddAuthorParams) middleware.Responder {
	author := models.Author{
		AuthorID: auth.nextId,
		Name:     request.Body.Name,
		Surname:  request.Body.Surname,
	}

	if author.Name == nil || *author.Name == "" {
		return &authors.GetAllAuthorsDefault{
			Payload: &models.Error{
				Code:    swag.Int64(http.StatusBadRequest),
				Message: swag.String("Bad Request"),
			},
		}
	}

	auth.authors[auth.nextId] = &author
	auth.nextId++
	return &authors.AddAuthorCreated{
		Payload: &authors.AddAuthorCreatedBody{
			ID: author.AuthorID,
		},
	}
}

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	implementation := Handler{}
	api := operations.NewExample2API(swaggerSpec)
	api.AuthorsGetAllAuthorsHandler = authors.GetAllAuthorsHandlerFunc(implementation.Handle) //назначаем хэндлер
	api.AuthorsAddAuthorHandler = authors.AddAuthorHandlerFunc(implementation.Handle2)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Book authors"
	parser.LongDescription = "This is server about authors"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	server.Port = 8080
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
