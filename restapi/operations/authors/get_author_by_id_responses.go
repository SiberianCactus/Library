// Code generated by go-swagger; DO NOT EDIT.

package authors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"books/models"
)

// GetAuthorByIDOKCode is the HTTP code returned for type GetAuthorByIDOK
const GetAuthorByIDOKCode int = 200

/*
GetAuthorByIDOK successful operation

swagger:response getAuthorByIdOK
*/
type GetAuthorByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Author `json:"body,omitempty"`
}

// NewGetAuthorByIDOK creates GetAuthorByIDOK with default headers values
func NewGetAuthorByIDOK() *GetAuthorByIDOK {

	return &GetAuthorByIDOK{}
}

// WithPayload adds the payload to the get author by Id o k response
func (o *GetAuthorByIDOK) WithPayload(payload *models.Author) *GetAuthorByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get author by Id o k response
func (o *GetAuthorByIDOK) SetPayload(payload *models.Author) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthorByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAuthorByIDBadRequestCode is the HTTP code returned for type GetAuthorByIDBadRequest
const GetAuthorByIDBadRequestCode int = 400

/*
GetAuthorByIDBadRequest Invalid ID supplied

swagger:response getAuthorByIdBadRequest
*/
type GetAuthorByIDBadRequest struct {
}

// NewGetAuthorByIDBadRequest creates GetAuthorByIDBadRequest with default headers values
func NewGetAuthorByIDBadRequest() *GetAuthorByIDBadRequest {

	return &GetAuthorByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetAuthorByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAuthorByIDNotFoundCode is the HTTP code returned for type GetAuthorByIDNotFound
const GetAuthorByIDNotFoundCode int = 404

/*
GetAuthorByIDNotFound Author not found

swagger:response getAuthorByIdNotFound
*/
type GetAuthorByIDNotFound struct {
}

// NewGetAuthorByIDNotFound creates GetAuthorByIDNotFound with default headers values
func NewGetAuthorByIDNotFound() *GetAuthorByIDNotFound {

	return &GetAuthorByIDNotFound{}
}

// WriteResponse to the client
func (o *GetAuthorByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
