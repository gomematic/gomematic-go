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

// ListTeamUsersReader is a Reader for the ListTeamUsers structure.
type ListTeamUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTeamUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewListTeamUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListTeamUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListTeamUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTeamUsersOK creates a ListTeamUsersOK with default headers values
func NewListTeamUsersOK() *ListTeamUsersOK {
	return &ListTeamUsersOK{}
}

/*ListTeamUsersOK handles this case with default header values.

A collection of team users
*/
type ListTeamUsersOK struct {
	Payload []*models.TeamUser
}

func (o *ListTeamUsersOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/users][%d] listTeamUsersOK  %+v", 200, o.Payload)
}

func (o *ListTeamUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamUsersForbidden creates a ListTeamUsersForbidden with default headers values
func NewListTeamUsersForbidden() *ListTeamUsersForbidden {
	return &ListTeamUsersForbidden{}
}

/*ListTeamUsersForbidden handles this case with default header values.

User is not authorized
*/
type ListTeamUsersForbidden struct {
	Payload *models.GeneralError
}

func (o *ListTeamUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/users][%d] listTeamUsersForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamUsersNotFound creates a ListTeamUsersNotFound with default headers values
func NewListTeamUsersNotFound() *ListTeamUsersNotFound {
	return &ListTeamUsersNotFound{}
}

/*ListTeamUsersNotFound handles this case with default header values.

Team not found
*/
type ListTeamUsersNotFound struct {
	Payload *models.GeneralError
}

func (o *ListTeamUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/users][%d] listTeamUsersNotFound  %+v", 404, o.Payload)
}

func (o *ListTeamUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamUsersDefault creates a ListTeamUsersDefault with default headers values
func NewListTeamUsersDefault(code int) *ListTeamUsersDefault {
	return &ListTeamUsersDefault{
		_statusCode: code,
	}
}

/*ListTeamUsersDefault handles this case with default header values.

Some error unrelated to the handler
*/
type ListTeamUsersDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the list team users default response
func (o *ListTeamUsersDefault) Code() int {
	return o._statusCode
}

func (o *ListTeamUsersDefault) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/users][%d] ListTeamUsers default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
