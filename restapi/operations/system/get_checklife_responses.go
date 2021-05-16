// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"customlogs/models"
)

// GetChecklifeOKCode is the HTTP code returned for type GetChecklifeOK
const GetChecklifeOKCode int = 200

/*GetChecklifeOK Ok

swagger:response getChecklifeOK
*/
type GetChecklifeOK struct {

	/*
	  In: Body
	*/
	Payload *models.System `json:"body,omitempty"`
}

// NewGetChecklifeOK creates GetChecklifeOK with default headers values
func NewGetChecklifeOK() *GetChecklifeOK {

	return &GetChecklifeOK{}
}

// WithPayload adds the payload to the get checklife o k response
func (o *GetChecklifeOK) WithPayload(payload *models.System) *GetChecklifeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get checklife o k response
func (o *GetChecklifeOK) SetPayload(payload *models.System) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChecklifeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChecklifeUnauthorizedCode is the HTTP code returned for type GetChecklifeUnauthorized
const GetChecklifeUnauthorizedCode int = 401

/*GetChecklifeUnauthorized Unauthorized


swagger:response getChecklifeUnauthorized
*/
type GetChecklifeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetChecklifeUnauthorized creates GetChecklifeUnauthorized with default headers values
func NewGetChecklifeUnauthorized() *GetChecklifeUnauthorized {

	return &GetChecklifeUnauthorized{}
}

// WithPayload adds the payload to the get checklife unauthorized response
func (o *GetChecklifeUnauthorized) WithPayload(payload models.Error) *GetChecklifeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get checklife unauthorized response
func (o *GetChecklifeUnauthorized) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChecklifeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetChecklifeForbiddenCode is the HTTP code returned for type GetChecklifeForbidden
const GetChecklifeForbiddenCode int = 403

/*GetChecklifeForbidden Forbidden


swagger:response getChecklifeForbidden
*/
type GetChecklifeForbidden struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetChecklifeForbidden creates GetChecklifeForbidden with default headers values
func NewGetChecklifeForbidden() *GetChecklifeForbidden {

	return &GetChecklifeForbidden{}
}

// WithPayload adds the payload to the get checklife forbidden response
func (o *GetChecklifeForbidden) WithPayload(payload models.Error) *GetChecklifeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get checklife forbidden response
func (o *GetChecklifeForbidden) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChecklifeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetChecklifeMethodNotAllowedCode is the HTTP code returned for type GetChecklifeMethodNotAllowed
const GetChecklifeMethodNotAllowedCode int = 405

/*GetChecklifeMethodNotAllowed Method Not Allowed


swagger:response getChecklifeMethodNotAllowed
*/
type GetChecklifeMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetChecklifeMethodNotAllowed creates GetChecklifeMethodNotAllowed with default headers values
func NewGetChecklifeMethodNotAllowed() *GetChecklifeMethodNotAllowed {

	return &GetChecklifeMethodNotAllowed{}
}

// WithPayload adds the payload to the get checklife method not allowed response
func (o *GetChecklifeMethodNotAllowed) WithPayload(payload models.Error) *GetChecklifeMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get checklife method not allowed response
func (o *GetChecklifeMethodNotAllowed) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChecklifeMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetChecklifeInternalServerErrorCode is the HTTP code returned for type GetChecklifeInternalServerError
const GetChecklifeInternalServerErrorCode int = 500

/*GetChecklifeInternalServerError Internal Server Error


swagger:response getChecklifeInternalServerError
*/
type GetChecklifeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetChecklifeInternalServerError creates GetChecklifeInternalServerError with default headers values
func NewGetChecklifeInternalServerError() *GetChecklifeInternalServerError {

	return &GetChecklifeInternalServerError{}
}

// WithPayload adds the payload to the get checklife internal server error response
func (o *GetChecklifeInternalServerError) WithPayload(payload models.Error) *GetChecklifeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get checklife internal server error response
func (o *GetChecklifeInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChecklifeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
