// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/AliShan302/Task_Schedular/gen/models"
)

// ListTasksOKCode is the HTTP code returned for type ListTasksOK
const ListTasksOKCode int = 200

/*
ListTasksOK tasks response

swagger:response listTasksOK
*/
type ListTasksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Task `json:"body,omitempty"`
}

// NewListTasksOK creates ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {

	return &ListTasksOK{}
}

// WithPayload adds the payload to the list tasks o k response
func (o *ListTasksOK) WithPayload(payload []*models.Task) *ListTasksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list tasks o k response
func (o *ListTasksOK) SetPayload(payload []*models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTasksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Task, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListTasksNotFoundCode is the HTTP code returned for type ListTasksNotFound
const ListTasksNotFoundCode int = 404

/*
ListTasksNotFound tasks not found

swagger:response listTasksNotFound
*/
type ListTasksNotFound struct {
}

// NewListTasksNotFound creates ListTasksNotFound with default headers values
func NewListTasksNotFound() *ListTasksNotFound {

	return &ListTasksNotFound{}
}

// WriteResponse to the client
func (o *ListTasksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ListTasksInternalServerErrorCode is the HTTP code returned for type ListTasksInternalServerError
const ListTasksInternalServerErrorCode int = 500

/*
ListTasksInternalServerError internal server error

swagger:response listTasksInternalServerError
*/
type ListTasksInternalServerError struct {
}

// NewListTasksInternalServerError creates ListTasksInternalServerError with default headers values
func NewListTasksInternalServerError() *ListTasksInternalServerError {

	return &ListTasksInternalServerError{}
}

// WriteResponse to the client
func (o *ListTasksInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
