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

// UpdateTaskReader is a Reader for the UpdateTask structure.
type UpdateTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /task] updateTask", response, response.Code())
	}
}

// NewUpdateTaskOK creates a UpdateTaskOK with default headers values
func NewUpdateTaskOK() *UpdateTaskOK {
	return &UpdateTaskOK{}
}

/*
UpdateTaskOK describes a response with status code 200, with default header values.

task updated
*/
type UpdateTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this update task o k response has a 2xx status code
func (o *UpdateTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update task o k response has a 3xx status code
func (o *UpdateTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task o k response has a 4xx status code
func (o *UpdateTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update task o k response has a 5xx status code
func (o *UpdateTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update task o k response a status code equal to that given
func (o *UpdateTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update task o k response
func (o *UpdateTaskOK) Code() int {
	return 200
}

func (o *UpdateTaskOK) Error() string {
	return fmt.Sprintf("[PUT /task][%d] updateTaskOK  %+v", 200, o.Payload)
}

func (o *UpdateTaskOK) String() string {
	return fmt.Sprintf("[PUT /task][%d] updateTaskOK  %+v", 200, o.Payload)
}

func (o *UpdateTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTaskNotFound creates a UpdateTaskNotFound with default headers values
func NewUpdateTaskNotFound() *UpdateTaskNotFound {
	return &UpdateTaskNotFound{}
}

/*
UpdateTaskNotFound describes a response with status code 404, with default header values.

tasks not found
*/
type UpdateTaskNotFound struct {
}

// IsSuccess returns true when this update task not found response has a 2xx status code
func (o *UpdateTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update task not found response has a 3xx status code
func (o *UpdateTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task not found response has a 4xx status code
func (o *UpdateTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update task not found response has a 5xx status code
func (o *UpdateTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update task not found response a status code equal to that given
func (o *UpdateTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update task not found response
func (o *UpdateTaskNotFound) Code() int {
	return 404
}

func (o *UpdateTaskNotFound) Error() string {
	return fmt.Sprintf("[PUT /task][%d] updateTaskNotFound ", 404)
}

func (o *UpdateTaskNotFound) String() string {
	return fmt.Sprintf("[PUT /task][%d] updateTaskNotFound ", 404)
}

func (o *UpdateTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTaskInternalServerError creates a UpdateTaskInternalServerError with default headers values
func NewUpdateTaskInternalServerError() *UpdateTaskInternalServerError {
	return &UpdateTaskInternalServerError{}
}

/*
UpdateTaskInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type UpdateTaskInternalServerError struct {
}

// IsSuccess returns true when this update task internal server error response has a 2xx status code
func (o *UpdateTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update task internal server error response has a 3xx status code
func (o *UpdateTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task internal server error response has a 4xx status code
func (o *UpdateTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update task internal server error response has a 5xx status code
func (o *UpdateTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update task internal server error response a status code equal to that given
func (o *UpdateTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update task internal server error response
func (o *UpdateTaskInternalServerError) Code() int {
	return 500
}

func (o *UpdateTaskInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /task][%d] updateTaskInternalServerError ", 500)
}

func (o *UpdateTaskInternalServerError) String() string {
	return fmt.Sprintf("[PUT /task][%d] updateTaskInternalServerError ", 500)
}

func (o *UpdateTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
