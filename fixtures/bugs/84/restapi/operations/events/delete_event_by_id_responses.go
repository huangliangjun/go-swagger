package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DeleteEventByIDNoContent Successful response

swagger:response deleteEventByIdNoContent
*/
type DeleteEventByIDNoContent struct {
}

// NewDeleteEventByIDNoContent creates DeleteEventByIDNoContent with default headers values
func NewDeleteEventByIDNoContent() *DeleteEventByIDNoContent {
	return &DeleteEventByIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteEventByIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteEventByIDDefault Generic Error

swagger:response deleteEventByIdDefault
*/
type DeleteEventByIDDefault struct {
	_statusCode int
}

// NewDeleteEventByIDDefault creates DeleteEventByIDDefault with default headers values
func NewDeleteEventByIDDefault(code int) *DeleteEventByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteEventByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete event by Id default response
func (o *DeleteEventByIDDefault) WithStatusCode(code int) *DeleteEventByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete event by Id default response
func (o *DeleteEventByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *DeleteEventByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
}
