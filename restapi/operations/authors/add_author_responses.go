// Code generated by go-swagger; DO NOT EDIT.

package authors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddAuthorCreatedCode is the HTTP code returned for type AddAuthorCreated
const AddAuthorCreatedCode int = 201

/*
AddAuthorCreated Created

swagger:response addAuthorCreated
*/
type AddAuthorCreated struct {

	/*
	  In: Body
	*/
	Payload *AddAuthorCreatedBody `json:"body,omitempty"`
}

// NewAddAuthorCreated creates AddAuthorCreated with default headers values
func NewAddAuthorCreated() *AddAuthorCreated {

	return &AddAuthorCreated{}
}

// WithPayload adds the payload to the add author created response
func (o *AddAuthorCreated) WithPayload(payload *AddAuthorCreatedBody) *AddAuthorCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add author created response
func (o *AddAuthorCreated) SetPayload(payload *AddAuthorCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAuthorCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAuthorMethodNotAllowedCode is the HTTP code returned for type AddAuthorMethodNotAllowed
const AddAuthorMethodNotAllowedCode int = 405

/*
AddAuthorMethodNotAllowed Invalid input

swagger:response addAuthorMethodNotAllowed
*/
type AddAuthorMethodNotAllowed struct {
}

// NewAddAuthorMethodNotAllowed creates AddAuthorMethodNotAllowed with default headers values
func NewAddAuthorMethodNotAllowed() *AddAuthorMethodNotAllowed {

	return &AddAuthorMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddAuthorMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
