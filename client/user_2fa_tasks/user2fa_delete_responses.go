// Code generated by go-swagger; DO NOT EDIT.

package user_2fa_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// User2faDeleteReader is a Reader for the User2faDelete structure.
type User2faDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *User2faDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUser2faDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUser2faDeleteOK creates a User2faDeleteOK with default headers values
func NewUser2faDeleteOK() *User2faDeleteOK {
	return &User2faDeleteOK{}
}

/*User2faDeleteOK handles this case with default header values.

successful operation
*/
type User2faDeleteOK struct {
	Payload *models.JSONResponseDataVoid
}

func (o *User2faDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /user/_2fa][%d] user2faDeleteOK  %+v", 200, o.Payload)
}

func (o *User2faDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
