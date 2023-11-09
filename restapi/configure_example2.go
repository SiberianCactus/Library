// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"sync"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"books/models"
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

	auth.authors = append(auth.authors, &author)
	auth.nextId++
	return &authors.AddAuthorCreated{
		Payload: &authors.AddAuthorCreatedBody{
			ID: author.AuthorID,
		},
	}
}

func (auth *Handler) GetAuthorIdHandler(request authors.GetAuthorByIDParams) middleware.Responder {
	auth.RLock()
	defer auth.RUnlock()

	for _, task := range auth.authors {
		if task.AuthorID == request.AuthorID {
			return &authors.GetAuthorByIDOK{
				Payload: &models.Author{
					AuthorID: request.AuthorID,
					Name:     task.Name,
					Surname:  task.Surname,
				},
			}
		}
	}
	return &authors.GetAuthorByIDNotFound{
		Payload: &authors.GetAuthorByIDNotFoundBody{
			Message: "Author is not found",
		},
	}
}

func (auth *Handler) DelAuthorHandler(request authors.DeleteAuthorParams) middleware.Responder {
	auth.Lock()
	defer auth.Unlock()

	for ind, task := range auth.authors {
		if task.AuthorID == request.AuthorID {
			auth.authors = append(auth.authors[:ind], auth.authors[ind+1:]...)
			break
		}
	}
	return &authors.DeleteAuthorCreated{
		Payload: &authors.DeleteAuthorCreatedBody{
			ID: request.AuthorID,
		},
	}
}

//go:generate swagger generate server --target ..\..\NewProj --name Example2 --spec ..\swagger.yaml --principal interface{}

func configureFlags(api *operations.Example2API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Example2API) http.Handler {
	// configure the api here\
	implementation := Handler{}
	api.AuthorsGetAllAuthorsHandler = authors.GetAllAuthorsHandlerFunc(implementation.Handle) //назначаем хэндлер
	api.AuthorsAddAuthorHandler = authors.AddAuthorHandlerFunc(implementation.AddAuthorHandler)
	api.AuthorsDeleteAuthorHandler = authors.DeleteAuthorHandlerFunc(implementation.DelAuthorHandler)
	api.AuthorsGetAuthorByIDHandler = authors.GetAuthorByIDHandlerFunc(implementation.GetAuthorIdHandler)
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AuthorsAddAuthorHandler == nil {
		api.AuthorsAddAuthorHandler = authors.AddAuthorHandlerFunc(func(params authors.AddAuthorParams) middleware.Responder {
			return middleware.NotImplemented("operation authors.AddAuthor has not yet been implemented")
		})
	}
	if api.AuthorsDeleteAuthorHandler == nil {
		api.AuthorsDeleteAuthorHandler = authors.DeleteAuthorHandlerFunc(func(params authors.DeleteAuthorParams) middleware.Responder {
			return middleware.NotImplemented("operation authors.DeleteAuthor has not yet been implemented")
		})
	}
	if api.AuthorsGetAllAuthorsHandler == nil {
		api.AuthorsGetAllAuthorsHandler = authors.GetAllAuthorsHandlerFunc(func(params authors.GetAllAuthorsParams) middleware.Responder {
			return middleware.NotImplemented("operation authors.GetAllAuthors has not yet been implemented")
		})
	}
	if api.AuthorsGetAuthorByIDHandler == nil {
		api.AuthorsGetAuthorByIDHandler = authors.GetAuthorByIDHandlerFunc(func(params authors.GetAuthorByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation authors.GetAuthorByID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
