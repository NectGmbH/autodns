// Code generated by go-swagger; DO NOT EDIT.

package id4me_identity_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// ConfirmId4MeIdentityReader is a Reader for the ConfirmId4MeIdentity structure.
type ConfirmId4MeIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfirmId4MeIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConfirmId4MeIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConfirmId4MeIdentityOK creates a ConfirmId4MeIdentityOK with default headers values
func NewConfirmId4MeIdentityOK() *ConfirmId4MeIdentityOK {
	return &ConfirmId4MeIdentityOK{}
}

/*ConfirmId4MeIdentityOK handles this case with default header values.

successful operation
*/
type ConfirmId4MeIdentityOK struct {
	Payload *models.JSONResponseDataId4MeIdentity
}

func (o *ConfirmId4MeIdentityOK) Error() string {
	return fmt.Sprintf("[PUT /id4MeIdentity/{name}/_confirm][%d] confirmId4MeIdentityOK  %+v", 200, o.Payload)
}

func (o *ConfirmId4MeIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataId4MeIdentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}