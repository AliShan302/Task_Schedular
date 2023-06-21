// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AliShan302/Task_Schedular/gen/models"
)

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /task] listTasks", response, response.Code())
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/* ListTasksOK describes a response with status code 200, with default header values.

tasks response
*/
type ListTasksOK struct {
	Payload []*models.Task
}

// IsSuccess returns true when this list tasks o k response has a 2xx status code
func (o *ListTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tasks o k response has a 3xx status code
func (o *ListTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks o k response has a 4xx status code
func (o *ListTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tasks o k response has a 5xx status code
func (o *ListTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks o k response a status code equal to that given
func (o *ListTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list tasks o k response
func (o *ListTasksOK) Code() int {
	return 200
}

func (o *ListTasksOK) Error() string {
	return fmt.Sprintf("[GET /task][%d] listTasksOK  %+v", 200, o.Payload)
}

func (o *ListTasksOK) String() string {
	return fmt.Sprintf("[GET /task][%d] listTasksOK  %+v", 200, o.Payload)
}

func (o *ListTasksOK) GetPayload() []*models.Task {
	return o.Payload
}

func (o *ListTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksNotFound creates a ListTasksNotFound with default headers values
func NewListTasksNotFound() *ListTasksNotFound {
	return &ListTasksNotFound{}
}

/* ListTasksNotFound describes a response with status code 404, with default header values.

tasks not found
*/
type ListTasksNotFound struct {
}

// IsSuccess returns true when this list tasks not found response has a 2xx status code
func (o *ListTasksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks not found response has a 3xx status code
func (o *ListTasksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks not found response has a 4xx status code
func (o *ListTasksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tasks not found response has a 5xx status code
func (o *ListTasksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks not found response a status code equal to that given
func (o *ListTasksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list tasks not found response
func (o *ListTasksNotFound) Code() int {
	return 404
}

func (o *ListTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /task][%d] listTasksNotFound ", 404)
}

func (o *ListTasksNotFound) String() string {
	return fmt.Sprintf("[GET /task][%d] listTasksNotFound ", 404)
}

func (o *ListTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTasksInternalServerError creates a ListTasksInternalServerError with default headers values
func NewListTasksInternalServerError() *ListTasksInternalServerError {
	return &ListTasksInternalServerError{}
}

/* ListTasksInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ListTasksInternalServerError struct {
}

// IsSuccess returns true when this list tasks internal server error response has a 2xx status code
func (o *ListTasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks internal server error response has a 3xx status code
func (o *ListTasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks internal server error response has a 4xx status code
func (o *ListTasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tasks internal server error response has a 5xx status code
func (o *ListTasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list tasks internal server error response a status code equal to that given
func (o *ListTasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list tasks internal server error response
func (o *ListTasksInternalServerError) Code() int {
	return 500
}

func (o *ListTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /task][%d] listTasksInternalServerError ", 500)
}

func (o *ListTasksInternalServerError) String() string {
	return fmt.Sprintf("[GET /task][%d] listTasksInternalServerError ", 500)
}

func (o *ListTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}