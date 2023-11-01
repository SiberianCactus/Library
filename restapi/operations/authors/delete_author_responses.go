// Code generated by go-swagger; DO NOT EDIT.

package authors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteAuthorBadRequestCode is the HTTP code returned for type DeleteAuthorBadRequest
const DeleteAuthorBadRequestCode int = 400

/*
DeleteAuthorBadRequest Invalid ID supplied

swagger:response deleteAuthorBadRequest
*/
type DeleteAuthorBadRequest struct {
}

// NewDeleteAuthorBadRequest creates DeleteAuthorBadRequest with default headers values
func NewDeleteAuthorBadRequest() *DeleteAuthorBadRequest {

	return &DeleteAuthorBadRequest{}
}

// WriteResponse to the client
func (o *DeleteAuthorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteAuthorNotFoundCode is the HTTP code returned for type DeleteAuthorNotFound
const DeleteAuthorNotFoundCode int = 404

/*
DeleteAuthorNotFound Author not found

swagger:response deleteAuthorNotFound
*/
type DeleteAuthorNotFound struct {
}

// NewDeleteAuthorNotFound creates DeleteAuthorNotFound with default headers values
func NewDeleteAuthorNotFound() *DeleteAuthorNotFound {

	return &DeleteAuthorNotFound{}
}

// WriteResponse to the client
func (o *DeleteAuthorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
