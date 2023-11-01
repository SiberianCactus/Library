// Code generated by go-swagger; DO NOT EDIT.

package authors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"books/models"
)

// GetAllAuthorsOKCode is the HTTP code returned for type GetAllAuthorsOK
const GetAllAuthorsOKCode int = 200

/*
GetAllAuthorsOK successful response

swagger:response getAllAuthorsOK
*/
type GetAllAuthorsOK struct {

	/*
	  In: Body
	*/
	Payload models.Authors `json:"body,omitempty"`
}

// NewGetAllAuthorsOK creates GetAllAuthorsOK with default headers values
func NewGetAllAuthorsOK() *GetAllAuthorsOK {

	return &GetAllAuthorsOK{}
}

// WithPayload adds the payload to the get all authors o k response
func (o *GetAllAuthorsOK) WithPayload(payload models.Authors) *GetAllAuthorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all authors o k response
func (o *GetAllAuthorsOK) SetPayload(payload models.Authors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllAuthorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Authors{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetAllAuthorsDefault Все нестандартное

swagger:response getAllAuthorsDefault
*/
type GetAllAuthorsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllAuthorsDefault creates GetAllAuthorsDefault with default headers values
func NewGetAllAuthorsDefault(code int) *GetAllAuthorsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAllAuthorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get all authors default response
func (o *GetAllAuthorsDefault) WithStatusCode(code int) *GetAllAuthorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all authors default response
func (o *GetAllAuthorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get all authors default response
func (o *GetAllAuthorsDefault) WithPayload(payload *models.Error) *GetAllAuthorsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all authors default response
func (o *GetAllAuthorsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllAuthorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}