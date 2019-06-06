// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gomematic/gomematic-go/models"
)

// UpdateTeamReader is a Reader for the UpdateTeam structure.
type UpdateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewUpdateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTeamOK creates a UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

/*UpdateTeamOK handles this case with default header values.

The updated team details
*/
type UpdateTeamOK struct {
	Payload *models.Team
}

func (o *UpdateTeamOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamForbidden creates a UpdateTeamForbidden with default headers values
func NewUpdateTeamForbidden() *UpdateTeamForbidden {
	return &UpdateTeamForbidden{}
}

/*UpdateTeamForbidden handles this case with default header values.

User is not authorized
*/
type UpdateTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *UpdateTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamPreconditionFailed creates a UpdateTeamPreconditionFailed with default headers values
func NewUpdateTeamPreconditionFailed() *UpdateTeamPreconditionFailed {
	return &UpdateTeamPreconditionFailed{}
}

/*UpdateTeamPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type UpdateTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *UpdateTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamUnprocessableEntity creates a UpdateTeamUnprocessableEntity with default headers values
func NewUpdateTeamUnprocessableEntity() *UpdateTeamUnprocessableEntity {
	return &UpdateTeamUnprocessableEntity{}
}

/*UpdateTeamUnprocessableEntity handles this case with default header values.

Failed to validate request
*/
type UpdateTeamUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamDefault creates a UpdateTeamDefault with default headers values
func NewUpdateTeamDefault(code int) *UpdateTeamDefault {
	return &UpdateTeamDefault{
		_statusCode: code,
	}
}

/*UpdateTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type UpdateTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the update team default response
func (o *UpdateTeamDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] UpdateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
