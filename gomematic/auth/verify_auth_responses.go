// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gomematic/gomematic-go/models"
)

// VerifyAuthReader is a Reader for the VerifyAuth structure.
type VerifyAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVerifyAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewVerifyAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVerifyAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVerifyAuthOK creates a VerifyAuthOK with default headers values
func NewVerifyAuthOK() *VerifyAuthOK {
	return &VerifyAuthOK{}
}

/*VerifyAuthOK handles this case with default header values.

Meta data of the provided token
*/
type VerifyAuthOK struct {
	Payload *models.AuthVerify
}

func (o *VerifyAuthOK) Error() string {
	return fmt.Sprintf("[GET /auth/verify/{token}][%d] verifyAuthOK  %+v", 200, o.Payload)
}

func (o *VerifyAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthVerify)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyAuthUnauthorized creates a VerifyAuthUnauthorized with default headers values
func NewVerifyAuthUnauthorized() *VerifyAuthUnauthorized {
	return &VerifyAuthUnauthorized{}
}

/*VerifyAuthUnauthorized handles this case with default header values.

Unauthorized if token is invalid
*/
type VerifyAuthUnauthorized struct {
	Payload *models.GeneralError
}

func (o *VerifyAuthUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/verify/{token}][%d] verifyAuthUnauthorized  %+v", 401, o.Payload)
}

func (o *VerifyAuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyAuthDefault creates a VerifyAuthDefault with default headers values
func NewVerifyAuthDefault(code int) *VerifyAuthDefault {
	return &VerifyAuthDefault{
		_statusCode: code,
	}
}

/*VerifyAuthDefault handles this case with default header values.

Some error unrelated to the handler
*/
type VerifyAuthDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the verify auth default response
func (o *VerifyAuthDefault) Code() int {
	return o._statusCode
}

func (o *VerifyAuthDefault) Error() string {
	return fmt.Sprintf("[GET /auth/verify/{token}][%d] VerifyAuth default  %+v", o._statusCode, o.Payload)
}

func (o *VerifyAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
